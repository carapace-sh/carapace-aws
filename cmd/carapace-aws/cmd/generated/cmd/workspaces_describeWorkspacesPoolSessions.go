package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspacesPoolSessionsCmd = &cobra.Command{
	Use:   "describe-workspaces-pool-sessions",
	Short: "Retrieves a list that describes the streaming sessions for a specified pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspacesPoolSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeWorkspacesPoolSessionsCmd).Standalone()

		workspaces_describeWorkspacesPoolSessionsCmd.Flags().String("limit", "", "The maximum size of each page of results.")
		workspaces_describeWorkspacesPoolSessionsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
		workspaces_describeWorkspacesPoolSessionsCmd.Flags().String("pool-id", "", "The identifier of the pool.")
		workspaces_describeWorkspacesPoolSessionsCmd.Flags().String("user-id", "", "The identifier of the user.")
		workspaces_describeWorkspacesPoolSessionsCmd.MarkFlagRequired("pool-id")
	})
	workspacesCmd.AddCommand(workspaces_describeWorkspacesPoolSessionsCmd)
}
