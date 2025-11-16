package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_restoreTableToPointInTimeCmd = &cobra.Command{
	Use:   "restore-table-to-point-in-time",
	Short: "Restores the specified table to the specified point in time within `EarliestRestorableDateTime` and `LatestRestorableDateTime`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_restoreTableToPointInTimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_restoreTableToPointInTimeCmd).Standalone()

		dynamodb_restoreTableToPointInTimeCmd.Flags().String("billing-mode-override", "", "The billing mode of the restored table.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("global-secondary-index-override", "", "List of global secondary indexes for the restored table.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("local-secondary-index-override", "", "List of local secondary indexes for the restored table.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("on-demand-throughput-override", "", "")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("provisioned-throughput-override", "", "Provisioned throughput settings for the restored table.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("restore-date-time", "", "Time in the past to restore the table to.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("source-table-arn", "", "The DynamoDB table that will be restored.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("source-table-name", "", "Name of the source table that is being restored.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("ssespecification-override", "", "The new server-side encryption settings for the restored table.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("target-table-name", "", "The name of the new table to which it must be restored to.")
		dynamodb_restoreTableToPointInTimeCmd.Flags().String("use-latest-restorable-time", "", "Restore the table to the latest possible time.")
		dynamodb_restoreTableToPointInTimeCmd.MarkFlagRequired("target-table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_restoreTableToPointInTimeCmd)
}
