package runner

import (
	"bytes"
	"fmt"

	"github.com/aryahadii/runandeh/configuration"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/sirupsen/logrus"
)

var (
	appContainers []string
)

// AppResponse contains output of app
type AppResponse struct {
	Out string
	Err string
}

// RunAppContainer creates a new container and executes app's binary inside
// container. It returns app's output as a string and any other things
// (ex. compilation error, container failure) as an error.
func RunAppContainer(request *RunRequest) (*AppResponse, error) {
	execFile, err := GetExecutableFilePath(request)
	if err != nil {
		return nil, err
	}

	// Get app image's name
	appImage := ""
	switch request.CodeLang {
	case LangCpp:
		appImage = configuration.GetInstance().GetString("docker.images.cpp")
	}

	// Create container

	containerID := fmt.Sprintf("app-container-%s-%d", request.CodeLang, request.ID)

	tmpDirMount := docker.HostMount{
		Source: execFile,
		Target: "/tmp/app",
		Type:   "bind",
		BindOptions: &docker.BindOptions{
			Propagation: "rprivate",
		},
	}

	hostConfig := &docker.HostConfig{
		Mounts: []docker.HostMount{
			tmpDirMount,
		},
	}

	containerConfig := &docker.Config{
		Image: appImage,
		Cmd:   []string{"/tmp/app"},
	}

	// Runandeh bridge config
	endpoints := map[string]*docker.EndpointConfig{
		configuration.GetInstance().GetString("docker.bridge-name"): &docker.EndpointConfig{
			NetworkID: bridgeNet.ID,
			Links: []string{
				fmt.Sprintf("db-container-%s-%d:db", request.DBMS, request.ID),
			},
		},
	}
	containerNetwork := &docker.NetworkingConfig{
		EndpointsConfig: endpoints,
	}

	containerOpts := docker.CreateContainerOptions{
		Name:             containerID,
		Config:           containerConfig,
		HostConfig:       hostConfig,
		NetworkingConfig: containerNetwork,
		Context:          ctx,
	}

	container, err := cli.CreateContainer(containerOpts)
	if err != nil {
		return nil, fmt.Errorf("can't create app's container [image:%s] (%v)", appImage, err)
	}

	// Start container
	if err := cli.StartContainer(container.ID, &docker.HostConfig{}); err != nil {
		return nil, fmt.Errorf("can't start app's container (%v)", err)
	}
	appContainers = append(appContainers, container.ID)
	logrus.WithField("container-id", container.ID).Info("app's container started")

	returnCode, err := cli.WaitContainer(containerID)
	if err != nil {
		return nil, fmt.Errorf("container exited (%d) with error: %v", returnCode, err)
	}

	var outStream, errStream bytes.Buffer
	logsOpts := docker.LogsOptions{
		Container:    container.ID,
		Context:      ctx,
		OutputStream: &outStream,
		ErrorStream:  &errStream,
		Stdout:       true,
		Stderr:       true,
	}
	cli.Logs(logsOpts)

	appResponse := &AppResponse{
		outStream.String(),
		errStream.String(),
	}

	return appResponse, nil
}
