package entity

type Account struct {
	Id            string
	UserId        string
	Type          string
	Balance       float64
	AccountNumber string
}

type AccountRequest struct {
	UserId  string  `validate:"required" json:"user_id"`
	Type    string  `validate:"required, oneof=debit credit loan" json:"type"`
	Balance float64 `json:"balance"`
}
