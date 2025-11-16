package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getLoggingConfigurationCmd = &cobra.Command{
	Use:   "get-logging-configuration",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getLoggingConfigurationCmd).Standalone()

		wafRegional_getLoggingConfigurationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the web ACL for which you want to get the [LoggingConfiguration]().")
		wafRegional_getLoggingConfigurationCmd.MarkFlagRequired("resource-arn")
	})
	wafRegionalCmd.AddCommand(wafRegional_getLoggingConfigurationCmd)
}
