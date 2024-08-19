package entity

type Transaction struct {
	Id          string
	FromAccount string
	ToAccount   string
	Amount      float64
}
