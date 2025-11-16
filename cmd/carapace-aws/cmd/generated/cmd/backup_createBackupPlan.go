package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createBackupPlanCmd = &cobra.Command{
	Use:   "create-backup-plan",
	Short: "Creates a backup plan using a backup plan name and backup rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createBackupPlanCmd).Standalone()

	backup_createBackupPlanCmd.Flags().String("backup-plan", "", "The body of a backup plan.")
	backup_createBackupPlanCmd.Flags().String("backup-plan-tags", "", "The tags to assign to the backup plan.")
	backup_createBackupPlanCmd.Flags().String("creator-request-id", "", "Identifies the request and allows failed requests to be retried without the risk of running the operation twice.")
	backup_createBackupPlanCmd.MarkFlagRequired("backup-plan")
	backupCmd.AddCommand(backup_createBackupPlanCmd)
}
