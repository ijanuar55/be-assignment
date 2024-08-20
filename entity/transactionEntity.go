package entity

type Transaction struct {
	Id          string  `json:"id"`
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
}
