package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_listCreatedArtifactsCmd = &cobra.Command{
	Use:   "list-created-artifacts",
	Short: "Lists the created artifacts attached to a given migration task in an update stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_listCreatedArtifactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_listCreatedArtifactsCmd).Standalone()

		mgh_listCreatedArtifactsCmd.Flags().String("max-results", "", "Maximum number of results to be returned per page.")
		mgh_listCreatedArtifactsCmd.Flags().String("migration-task-name", "", "Unique identifier that references the migration task.")
		mgh_listCreatedArtifactsCmd.Flags().String("next-token", "", "If a `NextToken` was returned by a previous call, there are more results available.")
		mgh_listCreatedArtifactsCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
		mgh_listCreatedArtifactsCmd.MarkFlagRequired("migration-task-name")
		mgh_listCreatedArtifactsCmd.MarkFlagRequired("progress-update-stream")
	})
	mghCmd.AddCommand(mgh_listCreatedArtifactsCmd)
}
