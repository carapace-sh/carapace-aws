package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deregisterWorkspaceDirectoryCmd = &cobra.Command{
	Use:   "deregister-workspace-directory",
	Short: "Deregisters the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deregisterWorkspaceDirectoryCmd).Standalone()

	workspaces_deregisterWorkspaceDirectoryCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	workspaces_deregisterWorkspaceDirectoryCmd.MarkFlagRequired("directory-id")
	workspacesCmd.AddCommand(workspaces_deregisterWorkspaceDirectoryCmd)
}
