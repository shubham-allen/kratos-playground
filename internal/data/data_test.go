package data

import (
	"github.com/Allen-Career-Institute/go-kratos-sample/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type MockDB struct {
	*gorm.DB
}

func TestNewData(t *testing.T) {
	c := &conf.Data{
		Database: &conf.Data_Database{
			Driver:                "mysql",
			Source:                "test_dsn",
			MaxIdleConns:          10,
			MaxOpenConns:          100,
			MaxConnLifetimeInMins: 30,
		},
	}

	logger := log.DefaultLogger

	_, _, err := NewData(c, logger)
	assert.Error(t, err)

}

func TestNewData_InvalidDriver(t *testing.T) {
	c := &conf.Data{
		Database: &conf.Data_Database{
			Driver: "invalid_driver",
		},
	}

	logger := log.DefaultLogger

	_, _, err := NewData(c, logger)
	assert.Nil(t, err)

}

func TestNewData_NoDBConnection(t *testing.T) {
	c := &conf.Data{
		Database: &conf.Data_Database{
			Driver: "mysql",
		},
	}

	logger := log.DefaultLogger

	_, _, err := NewData(c, logger)
	assert.Error(t, err)

}
