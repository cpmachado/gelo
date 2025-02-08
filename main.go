package main

import (
	"archive/zip"
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path"

	"go.cpmachado.pt/gelo/fide"
	"go.cpmachado.pt/gelo/internal/config"
)

var Version string = "0.1.0"

func init() {
	cfg := config.GetConfig()
	parseFlags(cfg)
	cfg.Apply()
	slog.Info("INIT", slog.Any("config", cfg))
}

func main() {
	cfg := config.GetConfig()

	if err := os.MkdirAll(cfg.Destination, os.ModeDir); err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	xurl, err := url.Parse(fide.XmlURL)
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}
	_, filename := path.Split(xurl.Path)
	file, err := os.OpenFile(path.Join(cfg.Destination, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}
	resp, err := http.Get(fide.XmlURL)
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}
	err = file.Close()
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}
	err = resp.Body.Close()
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("MAIN", slog.String("message", "Opening Zip"))
	archive, err := zip.OpenReader(path.Join(cfg.Destination, filename))
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("MAIN", slog.String("message", "Opening XML"))
	xmlFile, err := archive.File[0].Open()
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("MAIN", slog.String("message", "Decoding XML"))
	decoder := xml.NewDecoder(xmlFile)

	var players fide.Players
	if err = decoder.Decode(&players); err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	if err = xmlFile.Close(); err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	if err = archive.Close(); err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("MAIN", slog.String("message", "Encoding csv"))

	file, err = os.OpenFile(path.Join(cfg.Destination, "players.csv"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	w := csv.NewWriter(file)

	if err = w.Write(fide.PlayerCsvHeader); err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	for _, player := range players.Players {
		if err = w.Write(player.ToCsvRecord()); err != nil {
			slog.Error("MAIN", slog.Any("error", err))
			os.Exit(1)
		}
	}

	w.Flush()
	if err = w.Error(); err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}

	if err = file.Close(); err != nil {
		slog.Error("MAIN", slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("MAIN", slog.String("message", "Finish encoding csv"))
}

func parseFlags(cfg *config.Config) {
	var version bool
	flag.StringVar(&cfg.Destination, "d", cfg.Destination, "Destination directory for resources")
	flag.BoolVar(&version, "v", false, "Display version")
	flag.Parse()

	if version {
		displayVersion()
		os.Exit(0)
	}
}

func displayVersion() {
	fmt.Printf("gelo-%s Copyright (c) 2025 cpmachado", Version)
}
