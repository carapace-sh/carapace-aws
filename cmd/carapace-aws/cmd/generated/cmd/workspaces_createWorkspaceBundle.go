package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createWorkspaceBundleCmd = &cobra.Command{
	Use:   "create-workspace-bundle",
	Short: "Creates the specified WorkSpace bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createWorkspaceBundleCmd).Standalone()

	workspaces_createWorkspaceBundleCmd.Flags().String("bundle-description", "", "The description of the bundle.")
	workspaces_createWorkspaceBundleCmd.Flags().String("bundle-name", "", "The name of the bundle.")
	workspaces_createWorkspaceBundleCmd.Flags().String("compute-type", "", "")
	workspaces_createWorkspaceBundleCmd.Flags().String("image-id", "", "The identifier of the image that is used to create the bundle.")
	workspaces_createWorkspaceBundleCmd.Flags().String("root-storage", "", "")
	workspaces_createWorkspaceBundleCmd.Flags().String("tags", "", "The tags associated with the bundle.")
	workspaces_createWorkspaceBundleCmd.Flags().String("user-storage", "", "")
	workspaces_createWorkspaceBundleCmd.MarkFlagRequired("bundle-description")
	workspaces_createWorkspaceBundleCmd.MarkFlagRequired("bundle-name")
	workspaces_createWorkspaceBundleCmd.MarkFlagRequired("compute-type")
	workspaces_createWorkspaceBundleCmd.MarkFlagRequired("image-id")
	workspaces_createWorkspaceBundleCmd.MarkFlagRequired("user-storage")
	workspacesCmd.AddCommand(workspaces_createWorkspaceBundleCmd)
}
