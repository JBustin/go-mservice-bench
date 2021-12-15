package handlers

import (
	"fmt"

	"github.com/go-mservice-bench/lib/account"
	"github.com/go-mservice-bench/lib/injectors"
	"github.com/go-mservice-bench/lib/transaction"
)

func TransactionJob(d *injectors.DI, data string) error {
	t, err := transaction.FromString(data)
	d.Logger.Debug(fmt.Sprintf("%v", t))
	if err != nil {
		return err
	}

	var accounts []account.Account
	if err = d.Db.Account.FindByIds(&accounts, []uint{
		t.From, t.To,
	}, d.Config.DbLimit); err != nil {
		return err
	}

	// TODO
	// make two next operation inside a transaction

	// update "from" account
	if err = d.Db.Account.Update(&accounts[0], account.UpdateAccountInput{
		Id:     accounts[0].Id,
		Amount: accounts[0].Amount - t.Amount,
	}); err != nil {
		return err
	}

	// update "to" account
	if err = d.Db.Account.Update(&accounts[1], account.UpdateAccountInput{
		Id:     accounts[1].Id,
		Amount: accounts[1].Amount + t.Amount,
	}); err != nil {
		return err
	}

	return nil
}
