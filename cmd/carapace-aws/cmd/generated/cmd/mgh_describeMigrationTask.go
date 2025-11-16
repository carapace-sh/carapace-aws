package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_describeMigrationTaskCmd = &cobra.Command{
	Use:   "describe-migration-task",
	Short: "Retrieves a list of all attributes associated with a specific migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_describeMigrationTaskCmd).Standalone()

	mgh_describeMigrationTaskCmd.Flags().String("migration-task-name", "", "The identifier given to the MigrationTask.")
	mgh_describeMigrationTaskCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
	mgh_describeMigrationTaskCmd.MarkFlagRequired("migration-task-name")
	mgh_describeMigrationTaskCmd.MarkFlagRequired("progress-update-stream")
	mghCmd.AddCommand(mgh_describeMigrationTaskCmd)
}
