package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_disassociateBackupVaultMpaApprovalTeamCmd = &cobra.Command{
	Use:   "disassociate-backup-vault-mpa-approval-team",
	Short: "Removes the association between an MPA approval team and a backup vault, disabling the MPA approval workflow for restore operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_disassociateBackupVaultMpaApprovalTeamCmd).Standalone()

	backup_disassociateBackupVaultMpaApprovalTeamCmd.Flags().String("backup-vault-name", "", "The name of the backup vault from which to disassociate the MPA approval team.")
	backup_disassociateBackupVaultMpaApprovalTeamCmd.Flags().String("requester-comment", "", "An optional comment explaining the reason for disassociating the MPA approval team from the backup vault.")
	backup_disassociateBackupVaultMpaApprovalTeamCmd.MarkFlagRequired("backup-vault-name")
	backupCmd.AddCommand(backup_disassociateBackupVaultMpaApprovalTeamCmd)
}
