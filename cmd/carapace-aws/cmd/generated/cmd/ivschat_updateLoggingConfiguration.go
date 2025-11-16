package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_updateLoggingConfigurationCmd = &cobra.Command{
	Use:   "update-logging-configuration",
	Short: "Updates a specified logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_updateLoggingConfigurationCmd).Standalone()

	ivschat_updateLoggingConfigurationCmd.Flags().String("destination-configuration", "", "A complex type that contains a destination configuration for where chat content will be logged.")
	ivschat_updateLoggingConfigurationCmd.Flags().String("identifier", "", "Identifier of the logging configuration to be updated.")
	ivschat_updateLoggingConfigurationCmd.Flags().String("name", "", "Logging-configuration name.")
	ivschat_updateLoggingConfigurationCmd.MarkFlagRequired("identifier")
	ivschatCmd.AddCommand(ivschat_updateLoggingConfigurationCmd)
}
