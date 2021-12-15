package account

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestModel(t *testing.T) {
	var db *sql.DB

	db, mock, err := sqlmock.New() // mock sql.DB
	assert.Equal(t, nil, err, "should not raise an error")

	mock.ExpectQuery("select sqlite_version()").
		WithArgs().WillReturnRows(
		mock.NewRows([]string{"version"}).FromCSVString("3.8.10"),
	)

	gdb, err := gorm.Open(sqlite.Dialector{Conn: db}, &gorm.Config{})
	assert.Equal(t, nil, err, "should not raise an error")

	model := &Model{db: gdb}

	// Create
	i := CreateAccountInput{ClientId: 1, Amount: 100.0}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(i.ClientId, i.Amount).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	a, err := model.Create(i)

	assert.Equal(t, nil, err, "should not raise an error")
	assert.Equal(t, Account{Id: 0, ClientId: 1, Amount: 100.0}, a, "account should be created")

	// Update
	u := UpdateAccountInput{Id: a.Id, ClientId: a.ClientId, Amount: 200.0}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(u.ClientId, u.Amount, u.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = model.Update(&a, u)

	assert.Equal(t, nil, err, "should not raise an error")
	assert.Equal(t, Account{Id: 0, ClientId: 1, Amount: 200.0}, a, "account should be updated")

	// FindById
	a1 := Account{}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WithArgs(u.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "client_id", "amount"}).
				AddRow(a.Id, a.ClientId, a.Amount),
		)

	err = model.FindById(&a1, a.Id)

	assert.Equal(t, nil, err, "should not raise an error")
	assert.Equal(t, a, a1, "account should be find and populate")

	// Delete
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(u.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = model.Delete(a)

	assert.Equal(t, nil, err, "should not raise an error")

	// Check all mock expectations
	err = mock.ExpectationsWereMet()
	assert.Equal(t, nil, err, "all expectations should be checked")
}
