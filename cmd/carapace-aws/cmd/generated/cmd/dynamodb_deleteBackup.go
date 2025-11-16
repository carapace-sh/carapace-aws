package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_deleteBackupCmd = &cobra.Command{
	Use:   "delete-backup",
	Short: "Deletes an existing backup of a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_deleteBackupCmd).Standalone()

	dynamodb_deleteBackupCmd.Flags().String("backup-arn", "", "The ARN associated with the backup.")
	dynamodb_deleteBackupCmd.MarkFlagRequired("backup-arn")
	dynamodbCmd.AddCommand(dynamodb_deleteBackupCmd)
}
