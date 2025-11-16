package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listTrafficDistributionGroupsCmd = &cobra.Command{
	Use:   "list-traffic-distribution-groups",
	Short: "Lists traffic distribution groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listTrafficDistributionGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listTrafficDistributionGroupsCmd).Standalone()

		connect_listTrafficDistributionGroupsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listTrafficDistributionGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listTrafficDistributionGroupsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	connectCmd.AddCommand(connect_listTrafficDistributionGroupsCmd)
}
