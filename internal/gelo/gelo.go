package gelo

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go.cpmachado.pt/gelo/fide"
)

type (
	Player  = fide.FidePlayer
	Players = fide.FidePlayers
)

const (
	dst        = "output"
	xmlPlayers = "output/players_list_xml_foa.xml"
	csvPlayers = "output/players.csv"
	zipFile    = "output/players_list_xml.zip"
	fideUrl    = fide.XmlURL
)

func writeCsv(players Players) {
	n := len(players.Players)
	header := "id,name,country,sex,title,w_title,o_title,foa_title,rating,games,k,rapid_rating,rapid_games,rapid_k,blitz_rating,blitz_games,blitz_k,birthday,flag"

	w, _ := os.OpenFile(csvPlayers, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)

	if _, err := w.WriteString(header); err != nil {
		log.Fatal(err)
	}
	for i, p := range players.Players {
		if i > 0 && i%1000 == 0 {
			fmt.Printf("\r                              \rWritten %d/%d", i, n)
		}
		if _, err := w.WriteString("\n"); err != nil {
			log.Fatal(err)
		}
		if _, err := w.WriteString(p.String()); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := w.WriteString("\n"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\r                              \rWritten %d/%d\n", n, n)

	w.Close()
}

func readXml() Players {
	xmlFile, err := os.Open(xmlPlayers)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	var players Players

	err = xml.Unmarshal(byteValue, &players)
	if err != nil {
		log.Fatal(err)
	}
	return players
}

func retrieveListZip() {
	c := http.Client{}
	resp, err := c.Get(fideUrl)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	err = os.WriteFile(zipFile, body, 0o644)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
}

func unzipList() {
	archive, err := zip.OpenReader(zipFile)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		fmt.Println("unzipping file ", filePath)
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}

func ExtractAndGenerateCsv() {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		panic(err)
	}
	fmt.Printf("Retriving list from %s\n", fideUrl)
	retrieveListZip()
	fmt.Printf("Unzipping list\n")
	unzipList()
	fmt.Printf("Parsing Players\n")
	players := readXml()
	fmt.Printf("Writing csv to %s\n", csvPlayers)
	writeCsv(players)
}
