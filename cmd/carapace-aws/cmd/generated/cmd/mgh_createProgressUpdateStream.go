package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_createProgressUpdateStreamCmd = &cobra.Command{
	Use:   "create-progress-update-stream",
	Short: "Creates a progress update stream which is an AWS resource used for access control as well as a namespace for migration task names that is implicitly linked to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_createProgressUpdateStreamCmd).Standalone()

	mgh_createProgressUpdateStreamCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
	mgh_createProgressUpdateStreamCmd.Flags().String("progress-update-stream-name", "", "The name of the ProgressUpdateStream.")
	mgh_createProgressUpdateStreamCmd.MarkFlagRequired("progress-update-stream-name")
	mghCmd.AddCommand(mgh_createProgressUpdateStreamCmd)
}
