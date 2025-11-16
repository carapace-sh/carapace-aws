package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putSubscriptionFilterCmd = &cobra.Command{
	Use:   "put-subscription-filter",
	Short: "Creates or updates a subscription filter and associates it with the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putSubscriptionFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putSubscriptionFilterCmd).Standalone()

		logs_putSubscriptionFilterCmd.Flags().String("apply-on-transformed-logs", "", "This parameter is valid only for log groups that have an active log transformer.")
		logs_putSubscriptionFilterCmd.Flags().String("destination-arn", "", "The ARN of the destination to deliver matching log events to.")
		logs_putSubscriptionFilterCmd.Flags().String("distribution", "", "The method used to distribute log data to the destination.")
		logs_putSubscriptionFilterCmd.Flags().String("emit-system-fields", "", "A list of system fields to include in the log events sent to the subscription destination.")
		logs_putSubscriptionFilterCmd.Flags().String("field-selection-criteria", "", "A filter expression that specifies which log events should be processed by this subscription filter based on system fields such as source account and source region.")
		logs_putSubscriptionFilterCmd.Flags().String("filter-name", "", "A name for the subscription filter.")
		logs_putSubscriptionFilterCmd.Flags().String("filter-pattern", "", "A filter pattern for subscribing to a filtered stream of log events.")
		logs_putSubscriptionFilterCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_putSubscriptionFilterCmd.Flags().String("role-arn", "", "The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream.")
		logs_putSubscriptionFilterCmd.MarkFlagRequired("destination-arn")
		logs_putSubscriptionFilterCmd.MarkFlagRequired("filter-name")
		logs_putSubscriptionFilterCmd.MarkFlagRequired("filter-pattern")
		logs_putSubscriptionFilterCmd.MarkFlagRequired("log-group-name")
	})
	logsCmd.AddCommand(logs_putSubscriptionFilterCmd)
}
