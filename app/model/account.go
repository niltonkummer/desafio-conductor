package model

type Account struct {
	Model
	Status string `json:"status"`
}

type Accounts []Account
