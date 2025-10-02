# FPX (Federação Portuguesa de Xadrez)

Website: <https://fpx.pt/site>

Portuguese federation data:

- blitz: <http://elo.fpx.pt/elo_csv.php?lista=b>
- classic: <http://elo.fpx.pt/elo_csv.php>
- rapid: <http://elo.fpx.pt/elo_csv.php?lista=r>

It's a simple csv using ";" as separator and the fields.

It has some issues with encoding and random Carriage return in some records.

So far it appears, by order:
- FPX id
- Name
- Federation
- Sex
- Club ID
- Club Name
- Date of Birth: only year in format (2006-_-_)
- Number of games? (Still not sure)
- FIDE ID
- Rating
- Title
- Age Group
  + U08
  + U10
  + U12
  + U14
  + U16
  + U18
  + U20
  + Sen
  + S50
  + S65
- Flags: Only inactive?
- K factor? (appears so, but some are 0 and empty, are these unrated?)
