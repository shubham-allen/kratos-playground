// Code generated by finalunit, visit us at https://github.com/wimspaargaren/final-unit
package otel

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OtelSuite struct {
	suite.Suite
}

func (s *OtelSuite) TestInitOtelProviders0() {
	InitOtelProviders()

}

func TestOtelSuite(t *testing.T) {
	suite.Run(t, new(OtelSuite))
}
