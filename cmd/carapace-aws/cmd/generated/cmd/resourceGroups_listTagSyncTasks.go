package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_listTagSyncTasksCmd = &cobra.Command{
	Use:   "list-tag-sync-tasks",
	Short: "Returns a list of tag-sync tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_listTagSyncTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_listTagSyncTasksCmd).Standalone()

		resourceGroups_listTagSyncTasksCmd.Flags().String("filters", "", "The Amazon resource name (ARN) or name of the application group for which you want to return a list of tag-sync tasks.")
		resourceGroups_listTagSyncTasksCmd.Flags().String("max-results", "", "The maximum number of results to be included in the response.")
		resourceGroups_listTagSyncTasksCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_listTagSyncTasksCmd)
}
