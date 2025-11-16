package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteWorkspaceBundleCmd = &cobra.Command{
	Use:   "delete-workspace-bundle",
	Short: "Deletes the specified WorkSpace bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteWorkspaceBundleCmd).Standalone()

	workspaces_deleteWorkspaceBundleCmd.Flags().String("bundle-id", "", "The identifier of the bundle.")
	workspacesCmd.AddCommand(workspaces_deleteWorkspaceBundleCmd)
}
