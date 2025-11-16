package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_stopConfigurationRecorderCmd = &cobra.Command{
	Use:   "stop-configuration-recorder",
	Short: "Stops the customer managed configuration recorder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_stopConfigurationRecorderCmd).Standalone()

	config_stopConfigurationRecorderCmd.Flags().String("configuration-recorder-name", "", "The name of the customer managed configuration recorder that you want to stop.")
	config_stopConfigurationRecorderCmd.MarkFlagRequired("configuration-recorder-name")
	configCmd.AddCommand(config_stopConfigurationRecorderCmd)
}
