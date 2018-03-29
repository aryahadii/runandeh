package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

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
	defer runner.RemoveContainers()

	// mq.InitMessageQueue()
	// defer mq.Close()

	runner.InitRunner()
	reqTable := []*runner.RunRequest{
		&runner.RunRequest{
			ID:       0,
			Payload:  nil,
			CodeLang: runner.LangCpp,
			AppCode: `#include <iostream>
			int main() {
				std::cout << "SHIT" << std::endl;
				return 0;
			}`,
			DB:                 runner.DBPostrges,
			DBValidatorQueries: nil,
		},
	}

	for _, req := range reqTable {
		runner.Run(req)
	}
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
