package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createWorkspaceImageCmd = &cobra.Command{
	Use:   "create-workspace-image",
	Short: "Creates a new WorkSpace image from an existing WorkSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createWorkspaceImageCmd).Standalone()

	workspaces_createWorkspaceImageCmd.Flags().String("description", "", "The description of the new WorkSpace image.")
	workspaces_createWorkspaceImageCmd.Flags().String("name", "", "The name of the new WorkSpace image.")
	workspaces_createWorkspaceImageCmd.Flags().String("tags", "", "The tags that you want to add to the new WorkSpace image.")
	workspaces_createWorkspaceImageCmd.Flags().String("workspace-id", "", "The identifier of the source WorkSpace")
	workspaces_createWorkspaceImageCmd.MarkFlagRequired("description")
	workspaces_createWorkspaceImageCmd.MarkFlagRequired("name")
	workspaces_createWorkspaceImageCmd.MarkFlagRequired("workspace-id")
	workspacesCmd.AddCommand(workspaces_createWorkspaceImageCmd)
}
