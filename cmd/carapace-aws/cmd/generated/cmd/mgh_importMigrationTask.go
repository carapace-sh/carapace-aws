package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_importMigrationTaskCmd = &cobra.Command{
	Use:   "import-migration-task",
	Short: "Registers a new migration task which represents a server, database, etc., being migrated to AWS by a migration tool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_importMigrationTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_importMigrationTaskCmd).Standalone()

		mgh_importMigrationTaskCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
		mgh_importMigrationTaskCmd.Flags().String("migration-task-name", "", "Unique identifier that references the migration task.")
		mgh_importMigrationTaskCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
		mgh_importMigrationTaskCmd.MarkFlagRequired("migration-task-name")
		mgh_importMigrationTaskCmd.MarkFlagRequired("progress-update-stream")
	})
	mghCmd.AddCommand(mgh_importMigrationTaskCmd)
}
