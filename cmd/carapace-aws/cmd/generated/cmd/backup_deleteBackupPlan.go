package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteBackupPlanCmd = &cobra.Command{
	Use:   "delete-backup-plan",
	Short: "Deletes a backup plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteBackupPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_deleteBackupPlanCmd).Standalone()

		backup_deleteBackupPlanCmd.Flags().String("backup-plan-id", "", "Uniquely identifies a backup plan.")
		backup_deleteBackupPlanCmd.MarkFlagRequired("backup-plan-id")
	})
	backupCmd.AddCommand(backup_deleteBackupPlanCmd)
}
