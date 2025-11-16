package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRestoreTestingPlansCmd = &cobra.Command{
	Use:   "list-restore-testing-plans",
	Short: "Returns a list of restore testing plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRestoreTestingPlansCmd).Standalone()

	backup_listRestoreTestingPlansCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listRestoreTestingPlansCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backupCmd.AddCommand(backup_listRestoreTestingPlansCmd)
}
