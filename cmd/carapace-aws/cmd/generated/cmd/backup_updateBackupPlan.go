package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateBackupPlanCmd = &cobra.Command{
	Use:   "update-backup-plan",
	Short: "Updates the specified backup plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateBackupPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_updateBackupPlanCmd).Standalone()

		backup_updateBackupPlanCmd.Flags().String("backup-plan", "", "The body of a backup plan.")
		backup_updateBackupPlanCmd.Flags().String("backup-plan-id", "", "The ID of the backup plan.")
		backup_updateBackupPlanCmd.MarkFlagRequired("backup-plan")
		backup_updateBackupPlanCmd.MarkFlagRequired("backup-plan-id")
	})
	backupCmd.AddCommand(backup_updateBackupPlanCmd)
}
