// Coverage template
package biz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BaseRepositorySuite struct {
	suite.Suite
}

func TestBaseRepositorySuite(t *testing.T) {
	suite.Run(t, new(BaseRepositorySuite))
}
