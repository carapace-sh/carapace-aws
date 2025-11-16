package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_disassociateSourceResourceCmd = &cobra.Command{
	Use:   "disassociate-source-resource",
	Short: "Removes the association between a source resource and a migration task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_disassociateSourceResourceCmd).Standalone()

	mgh_disassociateSourceResourceCmd.Flags().String("dry-run", "", "This is an optional parameter that you can use to test whether the call will succeed.")
	mgh_disassociateSourceResourceCmd.Flags().String("migration-task-name", "", "A unique identifier that references the migration task.")
	mgh_disassociateSourceResourceCmd.Flags().String("progress-update-stream", "", "The name of the progress-update stream, which is used for access control as well as a namespace for migration-task names that is implicitly linked to your AWS account.")
	mgh_disassociateSourceResourceCmd.Flags().String("source-resource-name", "", "The name that was specified for the source resource.")
	mgh_disassociateSourceResourceCmd.MarkFlagRequired("migration-task-name")
	mgh_disassociateSourceResourceCmd.MarkFlagRequired("progress-update-stream")
	mgh_disassociateSourceResourceCmd.MarkFlagRequired("source-resource-name")
	mghCmd.AddCommand(mgh_disassociateSourceResourceCmd)
}
