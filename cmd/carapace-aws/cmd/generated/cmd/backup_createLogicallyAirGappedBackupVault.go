package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createLogicallyAirGappedBackupVaultCmd = &cobra.Command{
	Use:   "create-logically-air-gapped-backup-vault",
	Short: "Creates a logical container to where backups may be copied.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createLogicallyAirGappedBackupVaultCmd).Standalone()

	backup_createLogicallyAirGappedBackupVaultCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_createLogicallyAirGappedBackupVaultCmd.Flags().String("backup-vault-tags", "", "The tags to assign to the vault.")
	backup_createLogicallyAirGappedBackupVaultCmd.Flags().String("creator-request-id", "", "The ID of the creation request.")
	backup_createLogicallyAirGappedBackupVaultCmd.Flags().String("encryption-key-arn", "", "The ARN of the customer-managed KMS key to use for encrypting the logically air-gapped backup vault.")
	backup_createLogicallyAirGappedBackupVaultCmd.Flags().String("max-retention-days", "", "The maximum retention period that the vault retains its recovery points.")
	backup_createLogicallyAirGappedBackupVaultCmd.Flags().String("min-retention-days", "", "This setting specifies the minimum retention period that the vault retains its recovery points.")
	backup_createLogicallyAirGappedBackupVaultCmd.MarkFlagRequired("backup-vault-name")
	backup_createLogicallyAirGappedBackupVaultCmd.MarkFlagRequired("max-retention-days")
	backup_createLogicallyAirGappedBackupVaultCmd.MarkFlagRequired("min-retention-days")
	backupCmd.AddCommand(backup_createLogicallyAirGappedBackupVaultCmd)
}
