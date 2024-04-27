package integrationtest

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"
	"kingscomp/internal/repository/redis"
	"kingscomp/pkg/testhelper"
	"os"
	"testing"
)

var redisPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntgeration() {
		return
	}

	pool := testhelper.StartDockerPool()

	// setup redis container for testing
	redisResource := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest", func(res *dockertest.Resource) error {
		port := res.GetPort("6379/tcp")
		_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
		return err
	})

	redisPort = redisResource.GetPort("6379/tcp")

	// run tests
	exitCode := m.Run()
	if err := redisResource.Close(); err != nil {
		logrus.WithError(err).Fatalln("couldn't close redis resource")
	}
	os.Exit(exitCode)
}
