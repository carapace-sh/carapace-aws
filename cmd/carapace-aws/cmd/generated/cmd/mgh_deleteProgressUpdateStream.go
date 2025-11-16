package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_deleteProgressUpdateStreamCmd = &cobra.Command{
	Use:   "delete-progress-update-stream",
	Short: "Deletes a progress update stream, including all of its tasks, which was previously created as an AWS resource used for access control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_deleteProgressUpdateStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_deleteProgressUpdateStreamCmd).Standalone()

		mgh_deleteProgressUpdateStreamCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
		mgh_deleteProgressUpdateStreamCmd.Flags().String("progress-update-stream-name", "", "The name of the ProgressUpdateStream.")
		mgh_deleteProgressUpdateStreamCmd.MarkFlagRequired("progress-update-stream-name")
	})
	mghCmd.AddCommand(mgh_deleteProgressUpdateStreamCmd)
}
