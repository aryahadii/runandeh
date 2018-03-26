package main

import (
	"github.com/aryahadii/runandeh"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Runandeh's version",
	Run: func(cmd *cobra.Command, args []string) {
		logVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func logVersion() {
	logrus.Info("version   > ", runandeh.Version)
	logrus.Info("buildtime > ", runandeh.BuildTime)
	logrus.Info("commit    > ", runandeh.Commit)
}
