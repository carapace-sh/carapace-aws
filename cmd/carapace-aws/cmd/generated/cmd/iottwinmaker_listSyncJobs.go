package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listSyncJobsCmd = &cobra.Command{
	Use:   "list-sync-jobs",
	Short: "List all SyncJobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listSyncJobsCmd).Standalone()

	iottwinmaker_listSyncJobsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iottwinmaker_listSyncJobsCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iottwinmaker_listSyncJobsCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the sync job.")
	iottwinmaker_listSyncJobsCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_listSyncJobsCmd)
}
