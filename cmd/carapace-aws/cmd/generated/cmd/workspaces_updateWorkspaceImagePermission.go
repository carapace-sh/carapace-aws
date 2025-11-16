package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_updateWorkspaceImagePermissionCmd = &cobra.Command{
	Use:   "update-workspace-image-permission",
	Short: "Shares or unshares an image with one account in the same Amazon Web Services Region by specifying whether that account has permission to copy the image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_updateWorkspaceImagePermissionCmd).Standalone()

	workspaces_updateWorkspaceImagePermissionCmd.Flags().String("allow-copy-image", "", "The permission to copy the image.")
	workspaces_updateWorkspaceImagePermissionCmd.Flags().String("image-id", "", "The identifier of the image.")
	workspaces_updateWorkspaceImagePermissionCmd.Flags().String("shared-account-id", "", "The identifier of the Amazon Web Services account to share or unshare the image with.")
	workspaces_updateWorkspaceImagePermissionCmd.MarkFlagRequired("allow-copy-image")
	workspaces_updateWorkspaceImagePermissionCmd.MarkFlagRequired("image-id")
	workspaces_updateWorkspaceImagePermissionCmd.MarkFlagRequired("shared-account-id")
	workspacesCmd.AddCommand(workspaces_updateWorkspaceImagePermissionCmd)
}
