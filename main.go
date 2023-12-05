package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/izaakdale/simple-blockchain/internal/blockchain"
)

type Transaction struct {
	From   uuid.UUID `json:"from"`
	To     uuid.UUID `json:"to"`
	Amount int32     `json:"amount"`
}

func main() {
	fmt.Println("Building blockchain")

	nbc := blockchain.New()

	nt1 := Transaction{
		From:   uuid.New(),
		To:     uuid.New(),
		Amount: 250,
	}

	bytes1, _ := json.Marshal(nt1)
	nbc.AddBlock([]byte(bytes1))

	txCheck := Transaction{
		From:   nt1.From,
		To:     nt1.To,
		Amount: nt1.Amount,
	}

	txBytes, _ := json.Marshal(txCheck)
	fmt.Println("Transaction valid:", nbc.Verify(1, txBytes))
}
