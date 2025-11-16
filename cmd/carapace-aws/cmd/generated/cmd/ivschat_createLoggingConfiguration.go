package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_createLoggingConfigurationCmd = &cobra.Command{
	Use:   "create-logging-configuration",
	Short: "Creates a logging configuration that allows clients to store and record sent messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_createLoggingConfigurationCmd).Standalone()

	ivschat_createLoggingConfigurationCmd.Flags().String("destination-configuration", "", "A complex type that contains a destination configuration for where chat content will be logged.")
	ivschat_createLoggingConfigurationCmd.Flags().String("name", "", "Logging-configuration name.")
	ivschat_createLoggingConfigurationCmd.Flags().String("tags", "", "Tags to attach to the resource.")
	ivschat_createLoggingConfigurationCmd.MarkFlagRequired("destination-configuration")
	ivschatCmd.AddCommand(ivschat_createLoggingConfigurationCmd)
}
