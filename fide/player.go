/* Copyright © 2024 Carlos Pinto Machado<cpmachado@protonmail.com> */
package fide

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type FidePlayers struct {
	XMLName xml.Name     `xml:"playerslist"`
	Players []FidePlayer `xml:"player"`
}

type FidePlayer struct {
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

func (p FidePlayer) StringifiedRecords() []string {
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

func (p FidePlayer) String() string {
	return strings.Join(p.StringifiedRecords(), ";")
}
