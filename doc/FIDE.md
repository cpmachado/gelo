# FIDE

Currently there two main available formats.

- txt format: <https://ratings.fide.com/download/players_list.zip>
- xml format: <https://ratings.fide.com/download/players_list_xml.zip>

There's also a legacy format, which are per time control list.

XML schema is mapped in [lib/gelo.go](../lib/gelo.go)

TXT schema has this legend as below:

```txt
Legend:

STD/SRTNG - Standard rating
RPD/RRTNG - Rapid rating
BLZ/BRTNG - Blitz rating
SGM - number of STANDARD rated games in given period
RGM - number of RAPID rated games in given period
BGM - number of BLITZ rating games in given period
SK - STANDARD rating K factor
RK - RAPID rating K factor
BK - BLITZ rating K factor
B-day/BORN - year of birth of a player
ID NUMBER - identification number of a player within FIDE database
NAME - name of a player
TIT/TITL - title of a player (g - Grand Master, wg - Woman Grand Master, m - Interntional Master, wm - Woman International Master, f - FIDE Master, wf - Woman FIDE Master, c - Candidate Master, wc - Woman Candidate Master)
FED - Federation of a player
OTIT - Other titles of a player which may include (IA - International Arbiter, FA - FIDE Arbiter, NA - National Arbiter, IO - International Organizer, FT - FIDE Trainer, FST - FIDE Senior Trainer, DI - Developmental Instructor, NI - National Instructor)
FLAG - flag of inactivity (I - inactive, WI - woman inactive, w - woman)
SEX - sex of a player (M - male, F - female)
```

Reference: <https://ratings.fide.com/download_lists.phtml>
