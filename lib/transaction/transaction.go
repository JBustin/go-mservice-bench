package transaction

import (
	"fmt"
	"strconv"
	"strings"
)

const separator = ":"

type Transaction struct {
	From   uint    `json:"from"`
	To     uint    `json:"to"`
	Amount float64 `json:"amount"`
}

func NewTransaction(from, to uint, amount float64) Transaction {
	return Transaction{from, to, amount}
}

func FromString(t string) (Transaction, error) {
	sli := strings.Split(t, separator)
	if len(sli) < 2 {
		return Transaction{}, fmt.Errorf("invalid input")
	}
	from, err := strconv.Atoi(sli[0])
	if err != nil {
		return Transaction{}, err
	}
	to, err := strconv.Atoi(sli[1])
	if err != nil {
		return Transaction{}, err
	}
	amount, err := strconv.ParseFloat(sli[2], 64)
	if err != nil {
		return Transaction{}, err
	}
	return NewTransaction(uint(from), uint(to), amount), nil
}

func (t Transaction) String() string {
	return fmt.Sprintf("%v%v%v%v%v", t.From, separator, t.To, separator, t.Amount)
}
