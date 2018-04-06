package runner

import (
	"os"
	"time"

	"github.com/aryahadii/runandeh/configuration"
	docker "github.com/fsouza/go-dockerclient"

	"github.com/sirupsen/logrus"

	"golang.org/x/net/context"
)

var (
	ctx context.Context
	cli *docker.Client

	bridgeNet *docker.Network
)

// InitRunner creates docker client and pulls docker images
func InitRunner() {
	ctx = context.Background()

	// Init docker client
	var err error
	dockerEndpoint := "unix:///var/run/docker.sock"
	cli, err = docker.NewClient(dockerEndpoint)
	if err != nil {
		logrus.WithError(err).Fatal("can't init Docker client")
	}

	// Pull docker images
	if configuration.GetInstance().GetBool("docker.image-pull") {
		for name, ref := range configuration.GetInstance().GetStringMapString("docker.images") {
			logrus.Infof("pulling docker image: %s", name)

			// Create pull image configs
			pullOpts := docker.PullImageOptions{
				Repository: ref,
			}
			if configuration.GetInstance().GetBool("docker.image-pull-verbose") {
				pullOpts.OutputStream = os.Stdout
			}

			// Pull image
			if err := cli.PullImage(pullOpts, docker.AuthConfiguration{}); err != nil {
				logrus.WithError(err).Fatalf("can't pull docker images: %s", name)
			}
		}
	}

	if bridgeNet, err = createDockerNetwork(); err != nil {
		logrus.WithError(err).Fatal("can't create bridge net")
	}

	logrus.Info("runner initialized")
}

// Run creates a container and runs requested code inside the container
func Run(request *RunRequest) {
	if err := RunDBContainer(request); err != nil {
		logrus.WithField("ID", request.ID).WithError(err).Error("can't run db")
		return
	}

	// TODO: Implement better way to postpone app container's startup
	time.Sleep(5 * time.Second)

	response, err := RunAppContainer(request)
	if err != nil {
		logrus.WithField("ID", request.ID).WithError(err).Errorf("can't get executable")
		return
	}
	logrus.WithField("ID", request.ID).Debugf("response: %v", response)
}

// RemoveContainers remove all apps and dbs containers
func RemoveContainers() {
	containerIDs := append(appContainers, dbContainers...)
	logrus.Infof("remove containers: %v", containerIDs)
	for _, id := range containerIDs {
		cli.RemoveContainer(docker.RemoveContainerOptions{
			ID:            id,
			RemoveVolumes: true,
			Force:         true,
			Context:       ctx,
		})
	}
}

func createDockerNetwork() (*docker.Network, error) {
	// Search through networks
	nets, _ := cli.ListNetworks()
	for _, net := range nets {
		if net.Name == configuration.GetInstance().GetString("docker.bridge-name") {
			return &net, nil
		}
	}

	// Create new network
	netOpts := docker.CreateNetworkOptions{
		Name: configuration.GetInstance().GetString("docker.bridge-name"),
	}
	return cli.CreateNetwork(netOpts)
}
