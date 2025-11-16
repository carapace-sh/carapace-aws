package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_updateWorkspaceBundleCmd = &cobra.Command{
	Use:   "update-workspace-bundle",
	Short: "Updates a WorkSpace bundle with a new image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_updateWorkspaceBundleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_updateWorkspaceBundleCmd).Standalone()

		workspaces_updateWorkspaceBundleCmd.Flags().String("bundle-id", "", "The identifier of the bundle.")
		workspaces_updateWorkspaceBundleCmd.Flags().String("image-id", "", "The identifier of the image.")
	})
	workspacesCmd.AddCommand(workspaces_updateWorkspaceBundleCmd)
}
