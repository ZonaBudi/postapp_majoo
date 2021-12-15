package mysql_test

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"postapp/pkg/config"
	"postapp/pkg/mysql"
	"runtime"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MysqlTestSuite struct {
	suite.Suite
	resource *dockertest.Resource
	pool     *dockertest.Pool
}

func (suite *MysqlTestSuite) SetupTest() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	config := &config.EnvConfig{
		FileName: "config.test",
		Path:     basepath,
	}
	err := config.ReadConfig()
	if err != nil {
		suite.T().Error(err)
	}
	pool, err := dockertest.NewPool("")
	if err != nil {
		suite.T().Errorf("Could not connect to docker: %s", err)
	}
	suite.pool = pool
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_PASSWORD=password", "MYSQL_DATABASE=db", "MYSQL_USER=user", "MYSQL_ROOT_PASSWORD=password"})
	if err != nil {
		suite.T().Errorf("Could not start resource: %s", err)
	}
	viper.Set("mysql.port", resource.GetPort("3306/tcp"))
	suite.resource = resource
	pool.Retry(func() error {
		db, err := sql.Open("mysql", fmt.Sprintf("root:password@(localhost:%s)/db", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	})
}

func (suite *MysqlTestSuite) TestConnect() {
	_, err := mysql.Connect()
	assert.NoError(suite.T(), err)
	suite.pool.Purge(suite.resource)

}

func TestMysqlTestSuite(t *testing.T) {
	suite.Run(t, new(MysqlTestSuite))
}
