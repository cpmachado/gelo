package fpx

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"

	"github.com/icholy/replace"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

const (
	ClassicURL = "https://elo.fpx.pt/elo_csv.php"
	BlitzURL   = "http://elo.fpx.pt/elo_csv.php?lista=b"
	RapidURL   = "http://elo.fpx.pt/elo_csv.php?lista=r"
)

var CSVHeaders = []string{"id", "name", "country", "sex", "club_id", "club", "birthday", "unknown", "fideid", "rating", "title", "group", "flags", "k"}

type Player struct {
	ID       int        `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Country  string     `json:"country,omitempty"`
	Sex      string     `json:"sex,omitempty"`
	ClubID   int        `json:"club_id,omitempty"`
	Club     string     `json:"club,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Unknown  string     `json:"unknown,omitempty"`
	Fideid   string     `json:"fideid,omitempty"`
	Rating   int        `json:"rating,omitempty"`
	Title    string     `json:"title,omitempty"`
	Group    string     `json:"group,omitempty"`
	Flags    string     `json:"flags,omitempty"`
	K        int        `json:"k,omitempty"`
}

func ReadCSV(r io.Reader) ([][]string, error) {
	chain := transform.Chain(charmap.ISO8859_16.NewDecoder(), replace.String("\r\n", ""))
	t := transform.NewReader(r, chain)

	rcsv := csv.NewReader(t)
	rcsv.Comma = ';'

	records, err := rcsv.ReadAll()
	return records, err
}

func ParseRows(records [][]string) ([]Player, error) {
	var players []Player
	var err error

	for _, record := range records {
		p := Player{}
		p.ID, err = strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		p.Name = record[1]
		p.Country = record[2]
		p.Sex = record[3]
		p.ClubID, err = strconv.Atoi(record[4])
		if err != nil {
			return nil, err
		}
		p.Club = record[5]
		d, err := time.Parse("2006", record[6][:4])
		if err != nil {
			return nil, err
		}
		p.Birthday = &d
		players = append(players, p)
	}
	return players, nil
}
