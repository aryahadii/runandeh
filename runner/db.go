package runner

import (
	"fmt"

	"github.com/aryahadii/runandeh/configuration"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/sirupsen/logrus"
)

var (
	dbContainers []string
)

// RunDBContainer create a new container for requested database
func RunDBContainer(request *RunRequest) error {
	// Get DB image's name
	dbImage := ""
	switch request.DB {
	case DBPostrges:
		dbImage = configuration.GetInstance().GetString("docker.images.postgres")
	}

	// Create container
	containerID := fmt.Sprintf("db-container-%s-%d", request.DB, request.ID)
	container, err := cli.CreateContainer(docker.CreateContainerOptions{
		Name: containerID,
		Config: &docker.Config{
			Image: dbImage,
		},
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("can't create db container (%v)", err)
	}

	// Start container
	if err := cli.StartContainer(container.ID, &docker.HostConfig{}); err != nil {
		return fmt.Errorf("can't start db container (%v)", err)
	}
	dbContainers = append(dbContainers, container.ID)
	logrus.WithField("container-id", container.ID).Info("DB container started")

	return nil
}
