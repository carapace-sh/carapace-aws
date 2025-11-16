package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_putBackupVaultAccessPolicyCmd = &cobra.Command{
	Use:   "put-backup-vault-access-policy",
	Short: "Sets a resource-based policy that is used to manage access permissions on the target backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_putBackupVaultAccessPolicyCmd).Standalone()

	backup_putBackupVaultAccessPolicyCmd.Flags().String("backup-vault-name", "", "The name of a logical container where backups are stored.")
	backup_putBackupVaultAccessPolicyCmd.Flags().String("policy", "", "The backup vault access policy document in JSON format.")
	backup_putBackupVaultAccessPolicyCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_putBackupVaultAccessPolicyCmd)
}
