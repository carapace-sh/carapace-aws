package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_restoreTableFromBackupCmd = &cobra.Command{
	Use:   "restore-table-from-backup",
	Short: "Creates a new table from an existing backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_restoreTableFromBackupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_restoreTableFromBackupCmd).Standalone()

		dynamodb_restoreTableFromBackupCmd.Flags().String("backup-arn", "", "The Amazon Resource Name (ARN) associated with the backup.")
		dynamodb_restoreTableFromBackupCmd.Flags().String("billing-mode-override", "", "The billing mode of the restored table.")
		dynamodb_restoreTableFromBackupCmd.Flags().String("global-secondary-index-override", "", "List of global secondary indexes for the restored table.")
		dynamodb_restoreTableFromBackupCmd.Flags().String("local-secondary-index-override", "", "List of local secondary indexes for the restored table.")
		dynamodb_restoreTableFromBackupCmd.Flags().String("on-demand-throughput-override", "", "")
		dynamodb_restoreTableFromBackupCmd.Flags().String("provisioned-throughput-override", "", "Provisioned throughput settings for the restored table.")
		dynamodb_restoreTableFromBackupCmd.Flags().String("ssespecification-override", "", "The new server-side encryption settings for the restored table.")
		dynamodb_restoreTableFromBackupCmd.Flags().String("target-table-name", "", "The name of the new table to which the backup must be restored.")
		dynamodb_restoreTableFromBackupCmd.MarkFlagRequired("backup-arn")
		dynamodb_restoreTableFromBackupCmd.MarkFlagRequired("target-table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_restoreTableFromBackupCmd)
}
