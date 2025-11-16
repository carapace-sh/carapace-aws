package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getBackupPlanCmd = &cobra.Command{
	Use:   "get-backup-plan",
	Short: "Returns `BackupPlan` details for the specified `BackupPlanId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getBackupPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getBackupPlanCmd).Standalone()

		backup_getBackupPlanCmd.Flags().String("backup-plan-id", "", "Uniquely identifies a backup plan.")
		backup_getBackupPlanCmd.Flags().String("max-scheduled-runs-preview", "", "Number of future scheduled backup runs to preview.")
		backup_getBackupPlanCmd.Flags().String("version-id", "", "Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most 1,024 bytes long.")
		backup_getBackupPlanCmd.MarkFlagRequired("backup-plan-id")
	})
	backupCmd.AddCommand(backup_getBackupPlanCmd)
}
