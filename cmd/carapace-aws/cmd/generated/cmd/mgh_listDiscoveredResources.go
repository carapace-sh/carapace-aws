package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_listDiscoveredResourcesCmd = &cobra.Command{
	Use:   "list-discovered-resources",
	Short: "Lists discovered resources associated with the given `MigrationTask`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_listDiscoveredResourcesCmd).Standalone()

	mgh_listDiscoveredResourcesCmd.Flags().String("max-results", "", "The maximum number of results returned per page.")
	mgh_listDiscoveredResourcesCmd.Flags().String("migration-task-name", "", "The name of the MigrationTask.")
	mgh_listDiscoveredResourcesCmd.Flags().String("next-token", "", "If a `NextToken` was returned by a previous call, there are more results available.")
	mgh_listDiscoveredResourcesCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
	mgh_listDiscoveredResourcesCmd.MarkFlagRequired("migration-task-name")
	mgh_listDiscoveredResourcesCmd.MarkFlagRequired("progress-update-stream")
	mghCmd.AddCommand(mgh_listDiscoveredResourcesCmd)
}
