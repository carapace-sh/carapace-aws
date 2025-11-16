package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeBackupVaultCmd = &cobra.Command{
	Use:   "describe-backup-vault",
	Short: "Returns metadata about a backup vault specified by its name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeBackupVaultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_describeBackupVaultCmd).Standalone()

		backup_describeBackupVaultCmd.Flags().String("backup-vault-account-id", "", "The account ID of the specified backup vault.")
		backup_describeBackupVaultCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
		backup_describeBackupVaultCmd.MarkFlagRequired("backup-vault-name")
	})
	backupCmd.AddCommand(backup_describeBackupVaultCmd)
}
