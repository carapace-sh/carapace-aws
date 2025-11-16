package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspacesCmd = &cobra.Command{
	Use:   "describe-workspaces",
	Short: "Describes the specified WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspacesCmd).Standalone()

	workspaces_describeWorkspacesCmd.Flags().String("bundle-id", "", "The identifier of the bundle.")
	workspaces_describeWorkspacesCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	workspaces_describeWorkspacesCmd.Flags().String("limit", "", "The maximum number of items to return.")
	workspaces_describeWorkspacesCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeWorkspacesCmd.Flags().String("user-name", "", "The name of the directory user.")
	workspaces_describeWorkspacesCmd.Flags().String("workspace-ids", "", "The identifiers of the WorkSpaces.")
	workspaces_describeWorkspacesCmd.Flags().String("workspace-name", "", "The name of the user-decoupled WorkSpace.")
	workspacesCmd.AddCommand(workspaces_describeWorkspacesCmd)
}
