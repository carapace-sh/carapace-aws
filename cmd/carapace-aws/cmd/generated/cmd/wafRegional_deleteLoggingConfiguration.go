package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-logging-configuration",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteLoggingConfigurationCmd).Standalone()

		wafRegional_deleteLoggingConfigurationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the web ACL from which you want to delete the [LoggingConfiguration]().")
		wafRegional_deleteLoggingConfigurationCmd.MarkFlagRequired("resource-arn")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteLoggingConfigurationCmd)
}
