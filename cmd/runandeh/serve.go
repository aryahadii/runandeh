package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/aryahadii/runandeh/mq"
	"github.com/aryahadii/runandeh/runner"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve Runandeh",
		Run:   serve,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	logVersion()

	listenSignals()

	runner.InitRunner()

	mq.InitMessageQueue()
	mq.ListenToRunsQueue(func(request amqp.Delivery) {
		runner.Run(createRunRequest(request))
	})
}

func createRunRequest(request amqp.Delivery) *runner.RunRequest {
	runRequest := &runner.RunRequest{}
	json.Unmarshal(request.Body, runRequest)
	return runRequest
}

func listenSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGKILL)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		for sig := range c {
			logrus.Infof("signal '%s' received. trying to remove containers", sig.String())
			runner.RemoveContainers()
			os.Exit(0)
		}
	}()
}
