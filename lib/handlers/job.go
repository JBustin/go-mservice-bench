package handlers

import (
	"github.com/go-mservice-bench/lib/account"
	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/transaction"
)

func TransactionJob(db *db.DB, q *broker.Queue, data string) error {
	t, err := transaction.FromString(data)
	if err != nil {
		return err
	}

	var accounts []account.Account
	if err = db.Account.FindByIds(&accounts, []uint{
		t.From, t.To,
	}, db.Config.DbLimit); err != nil {
		return err
	}

	// TODO
	// make two next operation inside a transaction

	// update "from" account
	if err = db.Account.Update(&accounts[0], account.UpdateAccountInput{
		Amount: accounts[0].Amount - t.Amount,
	}); err != nil {
		return err
	}

	// update "to" account
	if err = db.Account.Update(&accounts[1], account.UpdateAccountInput{
		Amount: accounts[1].Amount + t.Amount,
	}); err != nil {
		return err
	}

	return nil
}
