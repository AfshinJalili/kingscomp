package testhelper

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
	"os"
)

type RetryFunc func(res *dockertest.Resource) error

func IsIntgeration() bool {
	return os.Getenv("TEST_INTEGRATION") == "true"
}

func StartDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.WithError(err).Fatalf("Could not cunstruct pool")
	}

	err = pool.Client.Ping()
	if err != nil {
		logrus.WithError(err).Fatalf("Could not connect to docker")
	}

	return pool
}

func StartDockerInstance(pool *dockertest.Pool, image, tag string, retryFunc RetryFunc, env ...string) *dockertest.Resource {

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: image,
		Tag:        tag,
		Env:        env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		logrus.WithError(err).Fatalf("Could not start resource")
	}

	if err := pool.Retry(func() error {
		return retryFunc(resource)
	}); err != nil {
		logrus.WithError(err).Fatalf("Could not connect to resource")
	}

	if err := resource.Expire(120); err != nil {
		logrus.WithError(err).Fatalf("Could not set expiration time")
	}

	return resource
}
