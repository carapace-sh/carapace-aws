package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_associateBackupVaultMpaApprovalTeamCmd = &cobra.Command{
	Use:   "associate-backup-vault-mpa-approval-team",
	Short: "Associates an MPA approval team with a backup vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_associateBackupVaultMpaApprovalTeamCmd).Standalone()

	backup_associateBackupVaultMpaApprovalTeamCmd.Flags().String("backup-vault-name", "", "The name of the backup vault to associate with the MPA approval team.")
	backup_associateBackupVaultMpaApprovalTeamCmd.Flags().String("mpa-approval-team-arn", "", "The Amazon Resource Name (ARN) of the MPA approval team to associate with the backup vault.")
	backup_associateBackupVaultMpaApprovalTeamCmd.Flags().String("requester-comment", "", "A comment provided by the requester explaining the association request.")
	backup_associateBackupVaultMpaApprovalTeamCmd.MarkFlagRequired("backup-vault-name")
	backup_associateBackupVaultMpaApprovalTeamCmd.MarkFlagRequired("mpa-approval-team-arn")
	backupCmd.AddCommand(backup_associateBackupVaultMpaApprovalTeamCmd)
}
