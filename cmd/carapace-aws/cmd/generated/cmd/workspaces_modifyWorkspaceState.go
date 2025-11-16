package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyWorkspaceStateCmd = &cobra.Command{
	Use:   "modify-workspace-state",
	Short: "Sets the state of the specified WorkSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyWorkspaceStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_modifyWorkspaceStateCmd).Standalone()

		workspaces_modifyWorkspaceStateCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
		workspaces_modifyWorkspaceStateCmd.Flags().String("workspace-state", "", "The WorkSpace state.")
		workspaces_modifyWorkspaceStateCmd.MarkFlagRequired("workspace-id")
		workspaces_modifyWorkspaceStateCmd.MarkFlagRequired("workspace-state")
	})
	workspacesCmd.AddCommand(workspaces_modifyWorkspaceStateCmd)
}
