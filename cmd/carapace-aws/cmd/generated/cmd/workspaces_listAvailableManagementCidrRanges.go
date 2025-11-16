package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_listAvailableManagementCidrRangesCmd = &cobra.Command{
	Use:   "list-available-management-cidr-ranges",
	Short: "Retrieves a list of IP address ranges, specified as IPv4 CIDR blocks, that you can use for the network management interface when you enable Bring Your Own License (BYOL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_listAvailableManagementCidrRangesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_listAvailableManagementCidrRangesCmd).Standalone()

		workspaces_listAvailableManagementCidrRangesCmd.Flags().String("management-cidr-range-constraint", "", "The IP address range to search.")
		workspaces_listAvailableManagementCidrRangesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		workspaces_listAvailableManagementCidrRangesCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
		workspaces_listAvailableManagementCidrRangesCmd.MarkFlagRequired("management-cidr-range-constraint")
	})
	workspacesCmd.AddCommand(workspaces_listAvailableManagementCidrRangesCmd)
}
