package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_putLoggingConfigurationCmd = &cobra.Command{
	Use:   "put-logging-configuration",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_putLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_putLoggingConfigurationCmd).Standalone()

		wafRegional_putLoggingConfigurationCmd.Flags().String("logging-configuration", "", "The Amazon Kinesis Data Firehose that contains the inspected traffic information, the redacted fields details, and the Amazon Resource Name (ARN) of the web ACL to monitor.")
		wafRegional_putLoggingConfigurationCmd.MarkFlagRequired("logging-configuration")
	})
	wafRegionalCmd.AddCommand(wafRegional_putLoggingConfigurationCmd)
}
