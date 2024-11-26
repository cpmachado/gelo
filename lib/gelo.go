/* Copyright © 2024 Carlos Pinto Machado<cpmachado@protonmail.com> */
package lib

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Player struct {
	XMLName      xml.Name `xml:"player"`
	Id           int      `xml:"fideid"`
	Name         string   `xml:"name"`
	Country      string   `xml:"country"`
	Sex          string   `xml:"sex"`
	Title        string   `xml:"title"`
	W_title      string   `xml:"w_title"`
	O_title      string   `xml:"o_title"`
	Foa_title    string   `xml:"foa_title"`
	Rating       int      `xml:"rating"`
	Games        int      `xml:"games"`
	K            int      `xml:"k"`
	Rapid_rating int      `xml:"rapid_rating"`
	Rapid_games  int      `xml:"rapid_games"`
	Rapid_k      int      `xml:"rapid_k"`
	Blitz_rating int      `xml:"blitz_rating"`
	Blitz_games  int      `xml:"blitz_games"`
	Blitz_k      int      `xml:"blitz_k"`
	Birthday     string   `xml:"birthday"`
	Flag         string   `xml:"flag"`
}

type Players struct {
	XMLName xml.Name `xml:"playerslist"`
	Players []Player `xml:"player"`
}

const (
	dst        = "output"
	xmlPlayers = "output/players_list_xml_foa.xml"
	csvPlayers = "output/players.csv"
	zipFile    = "output/players_list_xml.zip"
	fideUrl    = "https://ratings.fide.com/download/players_list_xml.zip"
)

func (p Player) StringifiedRecords() []string {
	return []string{
		strconv.Itoa(p.Id),
		p.Name,
		p.Country,
		p.Sex,
		p.Title,
		p.W_title,
		p.O_title,
		p.Foa_title,
		strconv.Itoa(p.Rating),
		strconv.Itoa(p.Games),
		strconv.Itoa(p.K),
		strconv.Itoa(p.Rapid_rating),
		strconv.Itoa(p.Rapid_games),
		strconv.Itoa(p.Rapid_k),
		strconv.Itoa(p.Blitz_rating),
		strconv.Itoa(p.Blitz_games),
		strconv.Itoa(p.Blitz_k),
		p.Birthday,
		p.Flag,
	}
}

func (p Player) String() string {
	return strings.Join(p.StringifiedRecords(), ";")
}

func writeCsv(players Players) {
	n := len(players.Players)
	header :=
		"id,name,country,sex,title,w_title,o_title,foa_title,rating,games,k,rapid_rating,rapid_games,rapid_k,blitz_rating,blitz_games,blitz_k,birthday,flag"

	w, _ := os.OpenFile(csvPlayers, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	w.WriteString(header)
	for i, p := range players.Players {
		if i > 0 && i%1000 == 0 {
			fmt.Printf("\r                              \rWritten %d/%d", i, n)
		}
		w.WriteString("\n")
		w.WriteString(p.String())
	}
	w.WriteString("\n")
	fmt.Printf("\r                              \rWritten %d/%d\n", n, n)

	w.Close()
}

func readXml() Players {
	xmlFile, err := os.Open(xmlPlayers)

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var players Players

	xml.Unmarshal(byteValue, &players)
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
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	err = ioutil.WriteFile(zipFile, body, 0777)

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
	fmt.Printf("Unzipping list to %s\n", xmlPlayers)
	unzipList()
	fmt.Printf("Parsing Players\n")
	players := readXml()
	fmt.Printf("Writing csv to %s\n", csvPlayers)
	writeCsv(players)
}
