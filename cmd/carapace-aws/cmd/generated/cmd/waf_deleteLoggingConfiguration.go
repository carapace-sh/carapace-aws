package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-logging-configuration",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteLoggingConfigurationCmd).Standalone()

	waf_deleteLoggingConfigurationCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the web ACL from which you want to delete the [LoggingConfiguration]().")
	waf_deleteLoggingConfigurationCmd.MarkFlagRequired("resource-arn")
	wafCmd.AddCommand(waf_deleteLoggingConfigurationCmd)
}
