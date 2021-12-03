package transaction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Transaction(t *testing.T) {
	transaction := NewTransaction(123, 234, 210.45)

	assert.Equal(t, 123, transaction.From, "should format the inputs")
	assert.Equal(t, "123:234:210.45", transaction.String(), "should stringified a transaction")

	newTransaction, err := FromString(transaction.String())

	assert.Equal(t, nil, err, "should not raise an error")
	assert.Equal(t, transaction, newTransaction, "should retrieve a serialized/unserialized transaction")

	_, err = FromString("unexpected string")

	assert.Equal(t, "invalid input", fmt.Sprintf("%v", err), "should raise an unserialized error")

	_, err = FromString("a:b:c")

	assert.Equal(t, "strconv.ParseFloat: parsing \"c\": invalid syntax", fmt.Sprintf("%v", err), "should raise a conversion error")
}
