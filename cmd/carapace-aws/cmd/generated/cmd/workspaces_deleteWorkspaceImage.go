package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteWorkspaceImageCmd = &cobra.Command{
	Use:   "delete-workspace-image",
	Short: "Deletes the specified image from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteWorkspaceImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_deleteWorkspaceImageCmd).Standalone()

		workspaces_deleteWorkspaceImageCmd.Flags().String("image-id", "", "The identifier of the image.")
		workspaces_deleteWorkspaceImageCmd.MarkFlagRequired("image-id")
	})
	workspacesCmd.AddCommand(workspaces_deleteWorkspaceImageCmd)
}
