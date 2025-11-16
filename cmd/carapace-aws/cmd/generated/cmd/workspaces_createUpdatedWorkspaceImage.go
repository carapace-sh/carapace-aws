package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createUpdatedWorkspaceImageCmd = &cobra.Command{
	Use:   "create-updated-workspace-image",
	Short: "Creates a new updated WorkSpace image based on the specified source image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createUpdatedWorkspaceImageCmd).Standalone()

	workspaces_createUpdatedWorkspaceImageCmd.Flags().String("description", "", "A description of whether updates for the WorkSpace image are available.")
	workspaces_createUpdatedWorkspaceImageCmd.Flags().String("name", "", "The name of the new updated WorkSpace image.")
	workspaces_createUpdatedWorkspaceImageCmd.Flags().String("source-image-id", "", "The identifier of the source WorkSpace image.")
	workspaces_createUpdatedWorkspaceImageCmd.Flags().String("tags", "", "The tags that you want to add to the new updated WorkSpace image.")
	workspaces_createUpdatedWorkspaceImageCmd.MarkFlagRequired("description")
	workspaces_createUpdatedWorkspaceImageCmd.MarkFlagRequired("name")
	workspaces_createUpdatedWorkspaceImageCmd.MarkFlagRequired("source-image-id")
	workspacesCmd.AddCommand(workspaces_createUpdatedWorkspaceImageCmd)
}
