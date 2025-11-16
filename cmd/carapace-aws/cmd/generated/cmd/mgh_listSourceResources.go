package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_listSourceResourcesCmd = &cobra.Command{
	Use:   "list-source-resources",
	Short: "Lists all the source resource that are associated with the specified `MigrationTaskName` and `ProgressUpdateStream`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_listSourceResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_listSourceResourcesCmd).Standalone()

		mgh_listSourceResourcesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		mgh_listSourceResourcesCmd.Flags().String("migration-task-name", "", "A unique identifier that references the migration task.")
		mgh_listSourceResourcesCmd.Flags().String("next-token", "", "If `NextToken` was returned by a previous call, there are more results available.")
		mgh_listSourceResourcesCmd.Flags().String("progress-update-stream", "", "The name of the progress-update stream, which is used for access control as well as a namespace for migration-task names that is implicitly linked to your AWS account.")
		mgh_listSourceResourcesCmd.MarkFlagRequired("migration-task-name")
		mgh_listSourceResourcesCmd.MarkFlagRequired("progress-update-stream")
	})
	mghCmd.AddCommand(mgh_listSourceResourcesCmd)
}
