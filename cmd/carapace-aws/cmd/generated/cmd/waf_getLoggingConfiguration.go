package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getLoggingConfigurationCmd = &cobra.Command{
	Use:   "get-logging-configuration",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_getLoggingConfigurationCmd).Standalone()

		waf_getLoggingConfigurationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the web ACL for which you want to get the [LoggingConfiguration]().")
		waf_getLoggingConfigurationCmd.MarkFlagRequired("resource-arn")
	})
	wafCmd.AddCommand(waf_getLoggingConfigurationCmd)
}
