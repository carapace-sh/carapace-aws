package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteRecoveryPointCmd = &cobra.Command{
	Use:   "delete-recovery-point",
	Short: "Deletes the recovery point specified by a recovery point ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteRecoveryPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_deleteRecoveryPointCmd).Standalone()

		backup_deleteRecoveryPointCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_deleteRecoveryPointCmd.Flags().String("recovery-point-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
		backup_deleteRecoveryPointCmd.MarkFlagRequired("backup-vault-name")
		backup_deleteRecoveryPointCmd.MarkFlagRequired("recovery-point-arn")
	})
	backupCmd.AddCommand(backup_deleteRecoveryPointCmd)
}
