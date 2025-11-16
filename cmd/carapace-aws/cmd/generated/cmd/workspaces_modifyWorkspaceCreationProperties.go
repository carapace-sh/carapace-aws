package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyWorkspaceCreationPropertiesCmd = &cobra.Command{
	Use:   "modify-workspace-creation-properties",
	Short: "Modify the default properties used to create WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyWorkspaceCreationPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_modifyWorkspaceCreationPropertiesCmd).Standalone()

		workspaces_modifyWorkspaceCreationPropertiesCmd.Flags().String("resource-id", "", "The identifier of the directory.")
		workspaces_modifyWorkspaceCreationPropertiesCmd.Flags().String("workspace-creation-properties", "", "The default properties for creating WorkSpaces.")
		workspaces_modifyWorkspaceCreationPropertiesCmd.MarkFlagRequired("resource-id")
		workspaces_modifyWorkspaceCreationPropertiesCmd.MarkFlagRequired("workspace-creation-properties")
	})
	workspacesCmd.AddCommand(workspaces_modifyWorkspaceCreationPropertiesCmd)
}
