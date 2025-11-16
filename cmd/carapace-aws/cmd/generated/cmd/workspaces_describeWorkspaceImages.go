package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspaceImagesCmd = &cobra.Command{
	Use:   "describe-workspace-images",
	Short: "Retrieves a list that describes one or more specified images, if the image identifiers are provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspaceImagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeWorkspaceImagesCmd).Standalone()

		workspaces_describeWorkspaceImagesCmd.Flags().String("image-ids", "", "The identifier of the image.")
		workspaces_describeWorkspaceImagesCmd.Flags().String("image-type", "", "The type (owned or shared) of the image.")
		workspaces_describeWorkspaceImagesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		workspaces_describeWorkspaceImagesCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	})
	workspacesCmd.AddCommand(workspaces_describeWorkspaceImagesCmd)
}
