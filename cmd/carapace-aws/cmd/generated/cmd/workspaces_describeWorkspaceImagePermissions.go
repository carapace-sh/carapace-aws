package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspaceImagePermissionsCmd = &cobra.Command{
	Use:   "describe-workspace-image-permissions",
	Short: "Describes the permissions that the owner of an image has granted to other Amazon Web Services accounts for an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspaceImagePermissionsCmd).Standalone()

	workspaces_describeWorkspaceImagePermissionsCmd.Flags().String("image-id", "", "The identifier of the image.")
	workspaces_describeWorkspaceImagePermissionsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	workspaces_describeWorkspaceImagePermissionsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeWorkspaceImagePermissionsCmd.MarkFlagRequired("image-id")
	workspacesCmd.AddCommand(workspaces_describeWorkspaceImagePermissionsCmd)
}
