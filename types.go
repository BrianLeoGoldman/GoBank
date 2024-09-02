package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	Firstname string `json:"firstname""`
	Lastname  string `json:"lastname"`
}

type Account struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstname string, lastname string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		Firstname: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(1000000)),
		// Balance is initialized as zero
		CreatedAt: time.Now().UTC(),
	}
}
