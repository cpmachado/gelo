package fide

import (
	_ "encoding/csv"
	"encoding/xml"
	"strconv"
)

type Players struct {
	XMLName xml.Name `xml:"playerslist"`
	Players []Player `xml:"player"`
}

type Player struct {
	XMLName      xml.Name `xml:"player" csv:"xml_name"`
	Id           int      `xml:"fideid" csv:"id"`
	Name         string   `xml:"name" csv:"name"`
	Country      string   `xml:"country" csv:"country"`
	Sex          string   `xml:"sex" csv:"sex"`
	Title        string   `xml:"title" csv:"title"`
	W_title      string   `xml:"w_title" csv:"w_title"`
	O_title      string   `xml:"o_title" csv:"o_title"`
	Foa_title    string   `xml:"foa_title" csv:"foa_title"`
	Rating       int      `xml:"rating" csv:"rating"`
	Games        int      `xml:"games" csv:"games"`
	K            int      `xml:"k" csv:"k"`
	Rapid_rating int      `xml:"rapid_rating" csv:"rapid_rating"`
	Rapid_games  int      `xml:"rapid_games" csv:"rapid_games"`
	Rapid_k      int      `xml:"rapid_k" csv:"rapid_k"`
	Blitz_rating int      `xml:"blitz_rating" csv:"blitz_rating"`
	Blitz_games  int      `xml:"blitz_games" csv:"blitz_games"`
	Blitz_k      int      `xml:"blitz_k" csv:"blitz_k"`
	Birthday     string   `xml:"birthday" csv:"birthday"`
	Flag         string   `xml:"flag" csv:"flag"`
}

var PlayerCsvHeader = []string{
	"id",
	"name",
	"country",
	"sex",
	"title",
	"w_title",
	"o_title",
	"foa_title",
	"rating",
	"games",
	"k",
	"rapid_rating",
	"rapid_games",
	"rapid_k",
	"blitz_rating",
	"blitz_games",
	"blitz_k",
	"birthday",
	"flag",
}

func (p *Player) ToCsvRecord() []string {
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
