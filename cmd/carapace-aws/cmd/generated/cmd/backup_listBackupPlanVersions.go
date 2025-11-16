package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listBackupPlanVersionsCmd = &cobra.Command{
	Use:   "list-backup-plan-versions",
	Short: "Returns version metadata of your backup plans, including Amazon Resource Names (ARNs), backup plan IDs, creation and deletion dates, plan names, and version IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listBackupPlanVersionsCmd).Standalone()

	backup_listBackupPlanVersionsCmd.Flags().String("backup-plan-id", "", "Uniquely identifies a backup plan.")
	backup_listBackupPlanVersionsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listBackupPlanVersionsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backup_listBackupPlanVersionsCmd.MarkFlagRequired("backup-plan-id")
	backupCmd.AddCommand(backup_listBackupPlanVersionsCmd)
}
