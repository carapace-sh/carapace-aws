package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getRestoreTestingInferredMetadataCmd = &cobra.Command{
	Use:   "get-restore-testing-inferred-metadata",
	Short: "This request returns the minimal required set of metadata needed to start a restore job with secure default settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getRestoreTestingInferredMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getRestoreTestingInferredMetadataCmd).Standalone()

		backup_getRestoreTestingInferredMetadataCmd.Flags().String("backup-vault-account-id", "", "The account ID of the specified backup vault.")
		backup_getRestoreTestingInferredMetadataCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_getRestoreTestingInferredMetadataCmd.Flags().String("recovery-point-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
		backup_getRestoreTestingInferredMetadataCmd.MarkFlagRequired("backup-vault-name")
		backup_getRestoreTestingInferredMetadataCmd.MarkFlagRequired("recovery-point-arn")
	})
	backupCmd.AddCommand(backup_getRestoreTestingInferredMetadataCmd)
}
