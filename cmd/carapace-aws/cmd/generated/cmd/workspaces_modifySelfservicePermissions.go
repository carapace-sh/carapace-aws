package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifySelfservicePermissionsCmd = &cobra.Command{
	Use:   "modify-selfservice-permissions",
	Short: "Modifies the self-service WorkSpace management capabilities for your users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifySelfservicePermissionsCmd).Standalone()

	workspaces_modifySelfservicePermissionsCmd.Flags().String("resource-id", "", "The identifier of the directory.")
	workspaces_modifySelfservicePermissionsCmd.Flags().String("selfservice-permissions", "", "The permissions to enable or disable self-service capabilities.")
	workspaces_modifySelfservicePermissionsCmd.MarkFlagRequired("resource-id")
	workspaces_modifySelfservicePermissionsCmd.MarkFlagRequired("selfservice-permissions")
	workspacesCmd.AddCommand(workspaces_modifySelfservicePermissionsCmd)
}
