package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspacesPoolsCmd = &cobra.Command{
	Use:   "describe-workspaces-pools",
	Short: "Describes the specified WorkSpaces Pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspacesPoolsCmd).Standalone()

	workspaces_describeWorkspacesPoolsCmd.Flags().String("filters", "", "The filter conditions for the WorkSpaces Pool to return.")
	workspaces_describeWorkspacesPoolsCmd.Flags().String("limit", "", "The maximum number of items to return.")
	workspaces_describeWorkspacesPoolsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeWorkspacesPoolsCmd.Flags().String("pool-ids", "", "The identifier of the WorkSpaces Pools.")
	workspacesCmd.AddCommand(workspaces_describeWorkspacesPoolsCmd)
}
