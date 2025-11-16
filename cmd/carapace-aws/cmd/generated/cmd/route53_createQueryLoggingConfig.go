package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createQueryLoggingConfigCmd = &cobra.Command{
	Use:   "create-query-logging-config",
	Short: "Creates a configuration for DNS query logging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createQueryLoggingConfigCmd).Standalone()

	route53_createQueryLoggingConfigCmd.Flags().String("cloud-watch-logs-log-group-arn", "", "The Amazon Resource Name (ARN) for the log group that you want to Amazon Route 53 to send query logs to.")
	route53_createQueryLoggingConfigCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone that you want to log queries for.")
	route53_createQueryLoggingConfigCmd.MarkFlagRequired("cloud-watch-logs-log-group-arn")
	route53_createQueryLoggingConfigCmd.MarkFlagRequired("hosted-zone-id")
	route53Cmd.AddCommand(route53_createQueryLoggingConfigCmd)
}
