package main

import (
	"github.com/aryahadii/runandeh/configuration"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "runandehd <subcommand>",
	Run: nil,
}

func init() {
	cobra.OnInitialize()

	configFilePath := "config.yaml"
	rootCmd.PersistentFlags().StringVarP(&configFilePath,
		"config", "c", configFilePath, "Path to the config file (eg ./config.yaml)")
	configuration.SetFilePath(configFilePath)

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if configuration.GetInstance().GetBool("debug") {
			configuration.SetDebugLogLevel(true)
		}
	}
}
