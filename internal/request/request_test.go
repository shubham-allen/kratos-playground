// Code generated by finalunit, visit us at https://github.com/wimspaargaren/final-unit
package request

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RequestSuite struct {
	suite.Suite
}

func (s *RequestSuite) TestPopulate0() {

	_ = Populate()

}

func (s *RequestSuite) TestextractRequestID1() {

	ctx := func() context.Context {
		return nil
	}()

	s.Panics(func() {
		extractRequestID(ctx)
	})

}

func TestRequestSuite(t *testing.T) {
	suite.Run(t, new(RequestSuite))
}