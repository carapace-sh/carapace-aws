package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createBackupSelectionCmd = &cobra.Command{
	Use:   "create-backup-selection",
	Short: "Creates a JSON document that specifies a set of resources to assign to a backup plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createBackupSelectionCmd).Standalone()

	backup_createBackupSelectionCmd.Flags().String("backup-plan-id", "", "The ID of the backup plan.")
	backup_createBackupSelectionCmd.Flags().String("backup-selection", "", "The body of a request to assign a set of resources to a backup plan.")
	backup_createBackupSelectionCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.")
	backup_createBackupSelectionCmd.MarkFlagRequired("backup-plan-id")
	backup_createBackupSelectionCmd.MarkFlagRequired("backup-selection")
	backupCmd.AddCommand(backup_createBackupSelectionCmd)
}
