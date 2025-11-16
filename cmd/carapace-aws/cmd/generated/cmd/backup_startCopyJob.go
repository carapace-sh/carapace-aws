package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_startCopyJobCmd = &cobra.Command{
	Use:   "start-copy-job",
	Short: "Starts a job to create a one-time copy of the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_startCopyJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_startCopyJobCmd).Standalone()

		backup_startCopyJobCmd.Flags().String("destination-backup-vault-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a destination backup vault to copy to; for example, `arn:aws:backup:us-east-1:123456789012:backup-vault:aBackupVault`.")
		backup_startCopyJobCmd.Flags().String("iam-role-arn", "", "Specifies the IAM role ARN used to copy the target recovery point; for example, `arn:aws:iam::123456789012:role/S3Access`.")
		backup_startCopyJobCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `StartCopyJob`.")
		backup_startCopyJobCmd.Flags().String("lifecycle", "", "")
		backup_startCopyJobCmd.Flags().String("recovery-point-arn", "", "An ARN that uniquely identifies a recovery point to use for the copy job; for example, arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.")
		backup_startCopyJobCmd.Flags().String("source-backup-vault-name", "", "The name of a logical source container where backups are stored.")
		backup_startCopyJobCmd.MarkFlagRequired("destination-backup-vault-arn")
		backup_startCopyJobCmd.MarkFlagRequired("iam-role-arn")
		backup_startCopyJobCmd.MarkFlagRequired("recovery-point-arn")
		backup_startCopyJobCmd.MarkFlagRequired("source-backup-vault-name")
	})
	backupCmd.AddCommand(backup_startCopyJobCmd)
}
