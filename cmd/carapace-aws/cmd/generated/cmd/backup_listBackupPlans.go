package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listBackupPlansCmd = &cobra.Command{
	Use:   "list-backup-plans",
	Short: "Lists the active backup plans for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listBackupPlansCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listBackupPlansCmd).Standalone()

		backup_listBackupPlansCmd.Flags().Bool("include-deleted", false, "A Boolean value with a default value of `FALSE` that returns deleted backup plans when set to `TRUE`.")
		backup_listBackupPlansCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
		backup_listBackupPlansCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
		backup_listBackupPlansCmd.Flags().Bool("no-include-deleted", false, "A Boolean value with a default value of `FALSE` that returns deleted backup plans when set to `TRUE`.")
		backup_listBackupPlansCmd.Flag("no-include-deleted").Hidden = true
	})
	backupCmd.AddCommand(backup_listBackupPlansCmd)
}
