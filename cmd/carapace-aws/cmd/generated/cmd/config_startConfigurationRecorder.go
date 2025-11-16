package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_startConfigurationRecorderCmd = &cobra.Command{
	Use:   "start-configuration-recorder",
	Short: "Starts the customer managed configuration recorder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_startConfigurationRecorderCmd).Standalone()

	config_startConfigurationRecorderCmd.Flags().String("configuration-recorder-name", "", "The name of the customer managed configuration recorder that you want to start.")
	config_startConfigurationRecorderCmd.MarkFlagRequired("configuration-recorder-name")
	configCmd.AddCommand(config_startConfigurationRecorderCmd)
}
