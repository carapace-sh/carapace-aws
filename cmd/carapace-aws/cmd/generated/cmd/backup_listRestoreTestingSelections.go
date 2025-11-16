package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRestoreTestingSelectionsCmd = &cobra.Command{
	Use:   "list-restore-testing-selections",
	Short: "Returns a list of restore testing selections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRestoreTestingSelectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listRestoreTestingSelectionsCmd).Standalone()

		backup_listRestoreTestingSelectionsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
		backup_listRestoreTestingSelectionsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
		backup_listRestoreTestingSelectionsCmd.Flags().String("restore-testing-plan-name", "", "Returns restore testing selections by the specified restore testing plan name.")
		backup_listRestoreTestingSelectionsCmd.MarkFlagRequired("restore-testing-plan-name")
	})
	backupCmd.AddCommand(backup_listRestoreTestingSelectionsCmd)
}
