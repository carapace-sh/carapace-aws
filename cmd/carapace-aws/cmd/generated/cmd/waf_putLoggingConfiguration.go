package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_putLoggingConfigurationCmd = &cobra.Command{
	Use:   "put-logging-configuration",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_putLoggingConfigurationCmd).Standalone()

	waf_putLoggingConfigurationCmd.Flags().String("logging-configuration", "", "The Amazon Kinesis Data Firehose that contains the inspected traffic information, the redacted fields details, and the Amazon Resource Name (ARN) of the web ACL to monitor.")
	waf_putLoggingConfigurationCmd.MarkFlagRequired("logging-configuration")
	wafCmd.AddCommand(waf_putLoggingConfigurationCmd)
}
