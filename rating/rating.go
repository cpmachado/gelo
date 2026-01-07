package rating

type RatingType int

const (
	EloRating RatingType = iota
)

type Rating interface {
	Type() RatingType
	Change(opponent Rating) (Rating, error)
	ChangeValue(opponent Rating) (float64, error)
}
