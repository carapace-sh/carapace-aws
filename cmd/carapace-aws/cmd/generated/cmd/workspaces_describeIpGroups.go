package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeIpGroupsCmd = &cobra.Command{
	Use:   "describe-ip-groups",
	Short: "Describes one or more of your IP access control groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeIpGroupsCmd).Standalone()

	workspaces_describeIpGroupsCmd.Flags().String("group-ids", "", "The identifiers of one or more IP access control groups.")
	workspaces_describeIpGroupsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	workspaces_describeIpGroupsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspacesCmd.AddCommand(workspaces_describeIpGroupsCmd)
}
