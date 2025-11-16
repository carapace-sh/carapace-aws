package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getLoggingConfigurationCmd = &cobra.Command{
	Use:   "get-logging-configuration",
	Short: "Returns the [LoggingConfiguration]() for the specified web ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getLoggingConfigurationCmd).Standalone()

	wafv2_getLoggingConfigurationCmd.Flags().String("log-scope", "", "The owner of the logging configuration, which must be set to `CUSTOMER` for the configurations that you manage.")
	wafv2_getLoggingConfigurationCmd.Flags().String("log-type", "", "Used to distinguish between various logging options.")
	wafv2_getLoggingConfigurationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the web ACL for which you want to get the [LoggingConfiguration]().")
	wafv2_getLoggingConfigurationCmd.MarkFlagRequired("resource-arn")
	wafv2Cmd.AddCommand(wafv2_getLoggingConfigurationCmd)
}
