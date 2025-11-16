package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_notifyMigrationTaskStateCmd = &cobra.Command{
	Use:   "notify-migration-task-state",
	Short: "Notifies Migration Hub of the current status, progress, or other detail regarding a migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_notifyMigrationTaskStateCmd).Standalone()

	mgh_notifyMigrationTaskStateCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
	mgh_notifyMigrationTaskStateCmd.Flags().String("migration-task-name", "", "Unique identifier that references the migration task.")
	mgh_notifyMigrationTaskStateCmd.Flags().String("next-update-seconds", "", "Number of seconds after the UpdateDateTime within which the Migration Hub can expect an update.")
	mgh_notifyMigrationTaskStateCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
	mgh_notifyMigrationTaskStateCmd.Flags().String("task", "", "Information about the task's progress and status.")
	mgh_notifyMigrationTaskStateCmd.Flags().String("update-date-time", "", "The timestamp when the task was gathered.")
	mgh_notifyMigrationTaskStateCmd.MarkFlagRequired("migration-task-name")
	mgh_notifyMigrationTaskStateCmd.MarkFlagRequired("next-update-seconds")
	mgh_notifyMigrationTaskStateCmd.MarkFlagRequired("progress-update-stream")
	mgh_notifyMigrationTaskStateCmd.MarkFlagRequired("task")
	mgh_notifyMigrationTaskStateCmd.MarkFlagRequired("update-date-time")
	mghCmd.AddCommand(mgh_notifyMigrationTaskStateCmd)
}
