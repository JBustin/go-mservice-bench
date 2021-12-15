package env

import (
	"testing"

	"github.com/go-mservice-bench/mocks"
	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	fs := mocks.NewFs()

	e, err := New(fs)

	assert.Equal(t, nil, err, "should not raise an error.")

	e.Data["VAR_FOO"] = "foo"
	e.Data["VAR_INT"] = "123"

	vStr, exists := e.Get("VAR_FOO")

	assert.Equal(t, true, exists, "should get a variable value")
	assert.Equal(t, "foo", vStr, "should get a variable value as string")

	vInt, exists := e.GetInt("VAR_INT")

	assert.Equal(t, true, exists, "should get a variable value")
	assert.Equal(t, 123, vInt, "should get a variable value as integer")

	_, exists = e.Get("NOT_EXIST")

	assert.Equal(t, false, exists, "should not get an unexisting string variable")

	_, exists = e.GetInt("NOT_EXIST")

	assert.Equal(t, false, exists, "should not get an unexisting integer variable")
}

func TestWalker(t *testing.T) {
	e := Env{
		Data: map[string]string{
			"VAR_FOO": "foo",
			"VAR_INT": "123",
		},
	}

	fn := e.walker()

	fn("VAR_HELLO=\"hello\"")

	vStr, exists := e.Get("VAR_HELLO")

	assert.Equal(t, true, exists, "should get a variable value")
	assert.Equal(t, "hello", vStr, "should get a variable value as string")

	fn("VAR_OLA=ola")

	vStr, exists = e.Get("VAR_OLA")

	assert.Equal(t, true, exists, "should get a variable value")
	assert.Equal(t, "ola", vStr, "should get a variable value as string")

	assert.Equal(t, 4, len(e.Data), "should contain a specific number of entries")

	fn("# COMMENT")
	fn("# VAR_HI=\"HI\"")
	fn("UNEXPECTED")

	assert.Equal(t, 4, len(e.Data), "should not import invalid cases")
}
