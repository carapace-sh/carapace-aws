package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_listMigrationTasksCmd = &cobra.Command{
	Use:   "list-migration-tasks",
	Short: "Lists all, or filtered by resource name, migration tasks associated with the user account making this call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_listMigrationTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_listMigrationTasksCmd).Standalone()

		mgh_listMigrationTasksCmd.Flags().String("max-results", "", "Value to specify how many results are returned per page.")
		mgh_listMigrationTasksCmd.Flags().String("next-token", "", "If a `NextToken` was returned by a previous call, there are more results available.")
		mgh_listMigrationTasksCmd.Flags().String("resource-name", "", "Filter migration tasks by discovered resource name.")
	})
	mghCmd.AddCommand(mgh_listMigrationTasksCmd)
}
