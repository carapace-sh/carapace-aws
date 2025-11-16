package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_listMigrationTaskUpdatesCmd = &cobra.Command{
	Use:   "list-migration-task-updates",
	Short: "This is a paginated API that returns all the migration-task states for the specified `MigrationTaskName` and `ProgressUpdateStream`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_listMigrationTaskUpdatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_listMigrationTaskUpdatesCmd).Standalone()

		mgh_listMigrationTaskUpdatesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		mgh_listMigrationTaskUpdatesCmd.Flags().String("migration-task-name", "", "A unique identifier that references the migration task.")
		mgh_listMigrationTaskUpdatesCmd.Flags().String("next-token", "", "If `NextToken` was returned by a previous call, there are more results available.")
		mgh_listMigrationTaskUpdatesCmd.Flags().String("progress-update-stream", "", "The name of the progress-update stream, which is used for access control as well as a namespace for migration-task names that is implicitly linked to your AWS account.")
		mgh_listMigrationTaskUpdatesCmd.MarkFlagRequired("migration-task-name")
		mgh_listMigrationTaskUpdatesCmd.MarkFlagRequired("progress-update-stream")
	})
	mghCmd.AddCommand(mgh_listMigrationTaskUpdatesCmd)
}
