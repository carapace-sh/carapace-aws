package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putConfigurationRecorderCmd = &cobra.Command{
	Use:   "put-configuration-recorder",
	Short: "Creates or updates the customer managed configuration recorder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putConfigurationRecorderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putConfigurationRecorderCmd).Standalone()

		config_putConfigurationRecorderCmd.Flags().String("configuration-recorder", "", "An object for the configuration recorder.")
		config_putConfigurationRecorderCmd.Flags().String("tags", "", "The tags for the customer managed configuration recorder.")
		config_putConfigurationRecorderCmd.MarkFlagRequired("configuration-recorder")
	})
	configCmd.AddCommand(config_putConfigurationRecorderCmd)
}
