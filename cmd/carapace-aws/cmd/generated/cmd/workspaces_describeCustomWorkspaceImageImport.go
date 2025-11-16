package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeCustomWorkspaceImageImportCmd = &cobra.Command{
	Use:   "describe-custom-workspace-image-import",
	Short: "Retrieves information about a WorkSpace BYOL image being imported via ImportCustomWorkspaceImage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeCustomWorkspaceImageImportCmd).Standalone()

	workspaces_describeCustomWorkspaceImageImportCmd.Flags().String("image-id", "", "The identifier of the WorkSpace image.")
	workspaces_describeCustomWorkspaceImageImportCmd.MarkFlagRequired("image-id")
	workspacesCmd.AddCommand(workspaces_describeCustomWorkspaceImageImportCmd)
}
