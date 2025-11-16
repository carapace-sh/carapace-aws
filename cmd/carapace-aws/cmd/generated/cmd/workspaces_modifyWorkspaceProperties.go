package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyWorkspacePropertiesCmd = &cobra.Command{
	Use:   "modify-workspace-properties",
	Short: "Modifies the specified WorkSpace properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyWorkspacePropertiesCmd).Standalone()

	workspaces_modifyWorkspacePropertiesCmd.Flags().String("data-replication", "", "Indicates the data replication status.")
	workspaces_modifyWorkspacePropertiesCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
	workspaces_modifyWorkspacePropertiesCmd.Flags().String("workspace-properties", "", "The properties of the WorkSpace.")
	workspaces_modifyWorkspacePropertiesCmd.MarkFlagRequired("workspace-id")
	workspacesCmd.AddCommand(workspaces_modifyWorkspacePropertiesCmd)
}
