package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_disassociateRecoveryPointCmd = &cobra.Command{
	Use:   "disassociate-recovery-point",
	Short: "Deletes the specified continuous backup recovery point from Backup and releases control of that continuous backup to the source service, such as Amazon RDS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_disassociateRecoveryPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_disassociateRecoveryPointCmd).Standalone()

		backup_disassociateRecoveryPointCmd.Flags().String("backup-vault-name", "", "The unique name of an Backup vault.")
		backup_disassociateRecoveryPointCmd.Flags().String("recovery-point-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies an Backup recovery point.")
		backup_disassociateRecoveryPointCmd.MarkFlagRequired("backup-vault-name")
		backup_disassociateRecoveryPointCmd.MarkFlagRequired("recovery-point-arn")
	})
	backupCmd.AddCommand(backup_disassociateRecoveryPointCmd)
}
