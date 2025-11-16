package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listSyncResourcesCmd = &cobra.Command{
	Use:   "list-sync-resources",
	Short: "Lists the sync resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listSyncResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_listSyncResourcesCmd).Standalone()

		iottwinmaker_listSyncResourcesCmd.Flags().String("filters", "", "A list of objects that filter the request.")
		iottwinmaker_listSyncResourcesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iottwinmaker_listSyncResourcesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iottwinmaker_listSyncResourcesCmd.Flags().String("sync-source", "", "The sync source.")
		iottwinmaker_listSyncResourcesCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the sync job.")
		iottwinmaker_listSyncResourcesCmd.MarkFlagRequired("sync-source")
		iottwinmaker_listSyncResourcesCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_listSyncResourcesCmd)
}
