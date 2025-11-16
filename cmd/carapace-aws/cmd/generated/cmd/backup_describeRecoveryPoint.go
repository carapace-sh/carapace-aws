package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeRecoveryPointCmd = &cobra.Command{
	Use:   "describe-recovery-point",
	Short: "Returns metadata associated with a recovery point, including ID, status, encryption, and lifecycle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeRecoveryPointCmd).Standalone()

	backup_describeRecoveryPointCmd.Flags().String("backup-vault-account-id", "", "The account ID of the specified backup vault.")
	backup_describeRecoveryPointCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_describeRecoveryPointCmd.Flags().String("recovery-point-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
	backup_describeRecoveryPointCmd.MarkFlagRequired("backup-vault-name")
	backup_describeRecoveryPointCmd.MarkFlagRequired("recovery-point-arn")
	backupCmd.AddCommand(backup_describeRecoveryPointCmd)
}
