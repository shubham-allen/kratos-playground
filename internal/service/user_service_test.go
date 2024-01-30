package service

import (
	test "github.com/Allen-Career-Institute/go-kratos-sample/tests/data/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateUser is running table tests with subtests.
// test cases are stored under test/data folder to improve readability of the test function.
func TestValidateUser(t *testing.T) {

	for _, tt := range test.TestCases {
		data := tt
		t.Run(data.Name, func(t *testing.T) {
			data.Mock()
			err := validateUser(data.Args.CreateUser)
			assert.Equal(t, err, data.ExpectedError)
		})
	}
}
