package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteBackupVaultCmd = &cobra.Command{
	Use:   "delete-backup-vault",
	Short: "Deletes the backup vault identified by its name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteBackupVaultCmd).Standalone()

	backup_deleteBackupVaultCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_deleteBackupVaultCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_deleteBackupVaultCmd)
}
