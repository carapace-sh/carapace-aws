package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_startRestoreJobCmd = &cobra.Command{
	Use:   "start-restore-job",
	Short: "Recovers the saved resource identified by an Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_startRestoreJobCmd).Standalone()

	backup_startRestoreJobCmd.Flags().String("copy-source-tags-to-restored-resource", "", "This is an optional parameter.")
	backup_startRestoreJobCmd.Flags().String("iam-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that Backup uses to create the target resource; for example: `arn:aws:iam::123456789012:role/S3Access`.")
	backup_startRestoreJobCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `StartRestoreJob`.")
	backup_startRestoreJobCmd.Flags().String("metadata", "", "A set of metadata key-value pairs.")
	backup_startRestoreJobCmd.Flags().String("recovery-point-arn", "", "An ARN that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
	backup_startRestoreJobCmd.Flags().String("resource-type", "", "Starts a job to restore a recovery point for one of the following resources:")
	backup_startRestoreJobCmd.MarkFlagRequired("metadata")
	backup_startRestoreJobCmd.MarkFlagRequired("recovery-point-arn")
	backupCmd.AddCommand(backup_startRestoreJobCmd)
}
