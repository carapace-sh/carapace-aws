package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_listBackupsCmd = &cobra.Command{
	Use:   "list-backups",
	Short: "List DynamoDB backups that are associated with an Amazon Web Services account and weren't made with Amazon Web Services Backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_listBackupsCmd).Standalone()

	dynamodb_listBackupsCmd.Flags().String("backup-type", "", "The backups from the table specified by `BackupType` are listed.")
	dynamodb_listBackupsCmd.Flags().String("exclusive-start-backup-arn", "", "`LastEvaluatedBackupArn` is the Amazon Resource Name (ARN) of the backup last evaluated when the current page of results was returned, inclusive of the current page of results.")
	dynamodb_listBackupsCmd.Flags().String("limit", "", "Maximum number of backups to return at once.")
	dynamodb_listBackupsCmd.Flags().String("table-name", "", "Lists the backups from the table specified in `TableName`.")
	dynamodb_listBackupsCmd.Flags().String("time-range-lower-bound", "", "Only backups created after this time are listed.")
	dynamodb_listBackupsCmd.Flags().String("time-range-upper-bound", "", "Only backups created before this time are listed.")
	dynamodbCmd.AddCommand(dynamodb_listBackupsCmd)
}
