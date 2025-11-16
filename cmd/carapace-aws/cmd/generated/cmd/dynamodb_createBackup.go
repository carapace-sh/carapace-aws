package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_createBackupCmd = &cobra.Command{
	Use:   "create-backup",
	Short: "Creates a backup for an existing table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_createBackupCmd).Standalone()

	dynamodb_createBackupCmd.Flags().String("backup-name", "", "Specified name for the backup.")
	dynamodb_createBackupCmd.Flags().String("table-name", "", "The name of the table.")
	dynamodb_createBackupCmd.MarkFlagRequired("backup-name")
	dynamodb_createBackupCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_createBackupCmd)
}
