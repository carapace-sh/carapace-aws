package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_copyWorkspaceImageCmd = &cobra.Command{
	Use:   "copy-workspace-image",
	Short: "Copies the specified image from the specified Region to the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_copyWorkspaceImageCmd).Standalone()

	workspaces_copyWorkspaceImageCmd.Flags().String("description", "", "A description of the image.")
	workspaces_copyWorkspaceImageCmd.Flags().String("name", "", "The name of the image.")
	workspaces_copyWorkspaceImageCmd.Flags().String("source-image-id", "", "The identifier of the source image.")
	workspaces_copyWorkspaceImageCmd.Flags().String("source-region", "", "The identifier of the source Region.")
	workspaces_copyWorkspaceImageCmd.Flags().String("tags", "", "The tags for the image.")
	workspaces_copyWorkspaceImageCmd.MarkFlagRequired("name")
	workspaces_copyWorkspaceImageCmd.MarkFlagRequired("source-image-id")
	workspaces_copyWorkspaceImageCmd.MarkFlagRequired("source-region")
	workspacesCmd.AddCommand(workspaces_copyWorkspaceImageCmd)
}
