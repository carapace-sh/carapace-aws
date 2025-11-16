package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeBackupCmd = &cobra.Command{
	Use:   "describe-backup",
	Short: "Describes an existing backup of a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeBackupCmd).Standalone()

	dynamodb_describeBackupCmd.Flags().String("backup-arn", "", "The Amazon Resource Name (ARN) associated with the backup.")
	dynamodb_describeBackupCmd.MarkFlagRequired("backup-arn")
	dynamodbCmd.AddCommand(dynamodb_describeBackupCmd)
}
