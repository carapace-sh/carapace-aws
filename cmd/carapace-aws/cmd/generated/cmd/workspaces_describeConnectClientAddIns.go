package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeConnectClientAddInsCmd = &cobra.Command{
	Use:   "describe-connect-client-add-ins",
	Short: "Retrieves a list of Amazon Connect client add-ins that have been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeConnectClientAddInsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_describeConnectClientAddInsCmd).Standalone()

		workspaces_describeConnectClientAddInsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		workspaces_describeConnectClientAddInsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
		workspaces_describeConnectClientAddInsCmd.Flags().String("resource-id", "", "The directory identifier for which the client add-in is configured.")
		workspaces_describeConnectClientAddInsCmd.MarkFlagRequired("resource-id")
	})
	workspacesCmd.AddCommand(workspaces_describeConnectClientAddInsCmd)
}
