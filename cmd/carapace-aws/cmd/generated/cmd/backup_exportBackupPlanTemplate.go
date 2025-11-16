package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_exportBackupPlanTemplateCmd = &cobra.Command{
	Use:   "export-backup-plan-template",
	Short: "Returns the backup plan that is specified by the plan ID as a backup template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_exportBackupPlanTemplateCmd).Standalone()

	backup_exportBackupPlanTemplateCmd.Flags().String("backup-plan-id", "", "Uniquely identifies a backup plan.")
	backup_exportBackupPlanTemplateCmd.MarkFlagRequired("backup-plan-id")
	backupCmd.AddCommand(backup_exportBackupPlanTemplateCmd)
}
