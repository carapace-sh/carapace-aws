package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deleteLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-logging-configuration",
	Short: "Deletes the [LoggingConfiguration]() from the specified web ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deleteLoggingConfigurationCmd).Standalone()

	wafv2_deleteLoggingConfigurationCmd.Flags().String("log-scope", "", "The owner of the logging configuration, which must be set to `CUSTOMER` for the configurations that you manage.")
	wafv2_deleteLoggingConfigurationCmd.Flags().String("log-type", "", "Used to distinguish between various logging options.")
	wafv2_deleteLoggingConfigurationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the web ACL from which you want to delete the [LoggingConfiguration]().")
	wafv2_deleteLoggingConfigurationCmd.MarkFlagRequired("resource-arn")
	wafv2Cmd.AddCommand(wafv2_deleteLoggingConfigurationCmd)
}
