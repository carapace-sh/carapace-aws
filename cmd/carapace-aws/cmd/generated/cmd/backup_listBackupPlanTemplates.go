package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listBackupPlanTemplatesCmd = &cobra.Command{
	Use:   "list-backup-plan-templates",
	Short: "Lists the backup plan templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listBackupPlanTemplatesCmd).Standalone()

	backup_listBackupPlanTemplatesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	backup_listBackupPlanTemplatesCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backupCmd.AddCommand(backup_listBackupPlanTemplatesCmd)
}
