package main

import (
	"github.com/aryahadii/runandeh/mq"
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

	mq.InitMessageQueue()
	defer mq.Close()
}
