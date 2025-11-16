package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getBackupPlanFromTemplateCmd = &cobra.Command{
	Use:   "get-backup-plan-from-template",
	Short: "Returns the template specified by its `templateId` as a backup plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getBackupPlanFromTemplateCmd).Standalone()

	backup_getBackupPlanFromTemplateCmd.Flags().String("backup-plan-template-id", "", "Uniquely identifies a stored backup plan template.")
	backup_getBackupPlanFromTemplateCmd.MarkFlagRequired("backup-plan-template-id")
	backupCmd.AddCommand(backup_getBackupPlanFromTemplateCmd)
}
