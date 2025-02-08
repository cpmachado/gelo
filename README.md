# gelo

[![build](https://github.com/cpmachado/gelo/actions/workflows/go.yml/badge.svg)](https://github.com/cpmachado/gelo/actions/workflows/go.yml)

Go ELO, is a simple program I started writting on the first day of WCC
2024(2024/11/25). After I got inspired by the brilliant performance of Ding
Liren.

Unfortunely, he lost, but did so as a champion.

## Currently
gelo simply retrieves the last xml list from FIDE and outputs a csv version of
it, which greatly reduces the size of the file and makes parsing easier.

## Of note
Given that the name field contains commas, the csv uses semi-colon(";") as a
separator.

## Usage
```sh
Usage of gelo:
  -d string
        Destination directory for resources (default "output")
  -v    Display version
```

It generates an "output" directory with the extracted resources and a csv
version of it.

## LICENSE

gelo is MIT Licensed as you can see in [LICENSE](LICENSE)
