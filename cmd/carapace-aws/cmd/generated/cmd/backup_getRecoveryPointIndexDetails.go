package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getRecoveryPointIndexDetailsCmd = &cobra.Command{
	Use:   "get-recovery-point-index-details",
	Short: "This operation returns the metadata and details specific to the backup index associated with the specified recovery point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getRecoveryPointIndexDetailsCmd).Standalone()

	backup_getRecoveryPointIndexDetailsCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_getRecoveryPointIndexDetailsCmd.Flags().String("recovery-point-arn", "", "An ARN that uniquely identifies a recovery point; for example, `arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.")
	backup_getRecoveryPointIndexDetailsCmd.MarkFlagRequired("backup-vault-name")
	backup_getRecoveryPointIndexDetailsCmd.MarkFlagRequired("recovery-point-arn")
	backupCmd.AddCommand(backup_getRecoveryPointIndexDetailsCmd)
}
