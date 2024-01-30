// Coverage template
package biz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BizSuite struct {
	suite.Suite
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(BizSuite))
}
