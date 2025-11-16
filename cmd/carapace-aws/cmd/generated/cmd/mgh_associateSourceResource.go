package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_associateSourceResourceCmd = &cobra.Command{
	Use:   "associate-source-resource",
	Short: "Associates a source resource with a migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_associateSourceResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_associateSourceResourceCmd).Standalone()

		mgh_associateSourceResourceCmd.Flags().String("dry-run", "", "This is an optional parameter that you can use to test whether the call will succeed.")
		mgh_associateSourceResourceCmd.Flags().String("migration-task-name", "", "A unique identifier that references the migration task.")
		mgh_associateSourceResourceCmd.Flags().String("progress-update-stream", "", "The name of the progress-update stream, which is used for access control as well as a namespace for migration-task names that is implicitly linked to your AWS account.")
		mgh_associateSourceResourceCmd.Flags().String("source-resource", "", "The source resource that you want to associate.")
		mgh_associateSourceResourceCmd.MarkFlagRequired("migration-task-name")
		mgh_associateSourceResourceCmd.MarkFlagRequired("progress-update-stream")
		mgh_associateSourceResourceCmd.MarkFlagRequired("source-resource")
	})
	mghCmd.AddCommand(mgh_associateSourceResourceCmd)
}
