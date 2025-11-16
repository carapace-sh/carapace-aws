package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspacesConnectionStatusCmd = &cobra.Command{
	Use:   "describe-workspaces-connection-status",
	Short: "Describes the connection status of the specified WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspacesConnectionStatusCmd).Standalone()

	workspaces_describeWorkspacesConnectionStatusCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeWorkspacesConnectionStatusCmd.Flags().String("workspace-ids", "", "The identifiers of the WorkSpaces.")
	workspacesCmd.AddCommand(workspaces_describeWorkspacesConnectionStatusCmd)
}
