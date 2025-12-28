package rating

type Elo struct{}

func (e *Elo) Type() RatingType {
	return EloRating
}

func (e *Elo) Change(opponent Rating) (Rating, error) {
	panic("not implemented") // TODO: Implement
}

func (e *Elo) ChangeValue(opponent Rating) (float64, error) {
	panic("not implemented") // TODO: Implement
}
