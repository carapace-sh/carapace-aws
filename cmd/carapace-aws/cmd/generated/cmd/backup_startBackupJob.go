package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_startBackupJobCmd = &cobra.Command{
	Use:   "start-backup-job",
	Short: "Starts an on-demand backup job for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_startBackupJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_startBackupJobCmd).Standalone()

		backup_startBackupJobCmd.Flags().String("backup-options", "", "The backup option for a selected resource.")
		backup_startBackupJobCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_startBackupJobCmd.Flags().String("complete-window-minutes", "", "A value in minutes during which a successfully started backup must complete, or else Backup will cancel the job.")
		backup_startBackupJobCmd.Flags().String("iam-role-arn", "", "Specifies the IAM role ARN used to create the target recovery point; for example, `arn:aws:iam::123456789012:role/S3Access`.")
		backup_startBackupJobCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `StartBackupJob`.")
		backup_startBackupJobCmd.Flags().String("index", "", "Include this parameter to enable index creation if your backup job has a resource type that supports backup indexes.")
		backup_startBackupJobCmd.Flags().String("lifecycle", "", "The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.")
		backup_startBackupJobCmd.Flags().String("recovery-point-tags", "", "The tags to assign to the resources.")
		backup_startBackupJobCmd.Flags().String("resource-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a resource.")
		backup_startBackupJobCmd.Flags().String("start-window-minutes", "", "A value in minutes after a backup is scheduled before a job will be canceled if it doesn't start successfully.")
		backup_startBackupJobCmd.MarkFlagRequired("backup-vault-name")
		backup_startBackupJobCmd.MarkFlagRequired("iam-role-arn")
		backup_startBackupJobCmd.MarkFlagRequired("resource-arn")
	})
	backupCmd.AddCommand(backup_startBackupJobCmd)
}
