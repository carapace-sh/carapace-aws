package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyWorkspaceAccessPropertiesCmd = &cobra.Command{
	Use:   "modify-workspace-access-properties",
	Short: "Specifies which devices and operating systems users can use to access their WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyWorkspaceAccessPropertiesCmd).Standalone()

	workspaces_modifyWorkspaceAccessPropertiesCmd.Flags().String("resource-id", "", "The identifier of the directory.")
	workspaces_modifyWorkspaceAccessPropertiesCmd.Flags().String("workspace-access-properties", "", "The device types and operating systems to enable or disable for access.")
	workspaces_modifyWorkspaceAccessPropertiesCmd.MarkFlagRequired("resource-id")
	workspaces_modifyWorkspaceAccessPropertiesCmd.MarkFlagRequired("workspace-access-properties")
	workspacesCmd.AddCommand(workspaces_modifyWorkspaceAccessPropertiesCmd)
}
