package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_createScheduledQueryCmd = &cobra.Command{
	Use:   "create-scheduled-query",
	Short: "Create a scheduled query that will be run on your behalf at the configured schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_createScheduledQueryCmd).Standalone()

	timestreamQuery_createScheduledQueryCmd.Flags().String("client-token", "", "Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("error-report-configuration", "", "Configuration for error reporting.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("kms-key-id", "", "The Amazon KMS key used to encrypt the scheduled query resource, at-rest.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("name", "", "Name of the scheduled query.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("notification-configuration", "", "Notification configuration for the scheduled query.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("query-string", "", "The query string to run.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("schedule-configuration", "", "The schedule configuration for the query.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("scheduled-query-execution-role-arn", "", "The ARN for the IAM role that Timestream will assume when running the scheduled query.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("tags", "", "A list of key-value pairs to label the scheduled query.")
	timestreamQuery_createScheduledQueryCmd.Flags().String("target-configuration", "", "Configuration used for writing the result of a query.")
	timestreamQuery_createScheduledQueryCmd.MarkFlagRequired("error-report-configuration")
	timestreamQuery_createScheduledQueryCmd.MarkFlagRequired("name")
	timestreamQuery_createScheduledQueryCmd.MarkFlagRequired("notification-configuration")
	timestreamQuery_createScheduledQueryCmd.MarkFlagRequired("query-string")
	timestreamQuery_createScheduledQueryCmd.MarkFlagRequired("schedule-configuration")
	timestreamQuery_createScheduledQueryCmd.MarkFlagRequired("scheduled-query-execution-role-arn")
	timestreamQueryCmd.AddCommand(timestreamQuery_createScheduledQueryCmd)
}
