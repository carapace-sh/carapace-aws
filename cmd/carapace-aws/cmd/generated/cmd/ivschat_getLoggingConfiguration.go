package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_getLoggingConfigurationCmd = &cobra.Command{
	Use:   "get-logging-configuration",
	Short: "Gets the specified logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_getLoggingConfigurationCmd).Standalone()

	ivschat_getLoggingConfigurationCmd.Flags().String("identifier", "", "Identifier of the logging configuration to be retrieved.")
	ivschat_getLoggingConfigurationCmd.MarkFlagRequired("identifier")
	ivschatCmd.AddCommand(ivschat_getLoggingConfigurationCmd)
}
