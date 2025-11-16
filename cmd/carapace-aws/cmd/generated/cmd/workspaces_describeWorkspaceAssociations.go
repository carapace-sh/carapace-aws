package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspaceAssociationsCmd = &cobra.Command{
	Use:   "describe-workspace-associations",
	Short: "Describes the associations betweens applications and the specified WorkSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspaceAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeWorkspaceAssociationsCmd).Standalone()

		workspaces_describeWorkspaceAssociationsCmd.Flags().String("associated-resource-types", "", "The resource types of the associated resources.")
		workspaces_describeWorkspaceAssociationsCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
		workspaces_describeWorkspaceAssociationsCmd.MarkFlagRequired("associated-resource-types")
		workspaces_describeWorkspaceAssociationsCmd.MarkFlagRequired("workspace-id")
	})
	workspacesCmd.AddCommand(workspaces_describeWorkspaceAssociationsCmd)
}
