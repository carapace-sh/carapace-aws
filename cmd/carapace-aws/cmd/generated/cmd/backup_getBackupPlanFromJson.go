package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getBackupPlanFromJsonCmd = &cobra.Command{
	Use:   "get-backup-plan-from-json",
	Short: "Returns a valid JSON document specifying a backup plan or an error.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getBackupPlanFromJsonCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getBackupPlanFromJsonCmd).Standalone()

		backup_getBackupPlanFromJsonCmd.Flags().String("backup-plan-template-json", "", "A customer-supplied backup plan document in JSON format.")
		backup_getBackupPlanFromJsonCmd.MarkFlagRequired("backup-plan-template-json")
	})
	backupCmd.AddCommand(backup_getBackupPlanFromJsonCmd)
}
