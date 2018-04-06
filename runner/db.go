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
	switch request.DBMS {
	case DBPostgres:
		dbImage = configuration.GetInstance().GetString("docker.images.postgres")
	case DBMySQL:
		dbImage = configuration.GetInstance().GetString("docker.images.mysql")
	}

	// Create container

	containerID := fmt.Sprintf("db-container-%s-%d", request.DBMS, request.ID)

	containerConfig := &docker.Config{
		Image: dbImage,
		Env: []string{
			"MYSQL_ALLOW_EMPTY_PASSWORD=yes",
			"MYSQL_DATABASE=test",
		},
	}

	// Runandeh bridge config
	endpoints := map[string]*docker.EndpointConfig{
		configuration.GetInstance().GetString("docker.bridge-name"): &docker.EndpointConfig{
			NetworkID: bridgeNet.ID,
			Aliases: []string{
				containerID,
			},
		},
	}
	containerNetwork := &docker.NetworkingConfig{
		EndpointsConfig: endpoints,
	}

	container, err := cli.CreateContainer(docker.CreateContainerOptions{
		Name:             containerID,
		Config:           containerConfig,
		NetworkingConfig: containerNetwork,
		Context:          ctx,
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
