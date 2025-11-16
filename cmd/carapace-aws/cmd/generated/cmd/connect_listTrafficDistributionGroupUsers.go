package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listTrafficDistributionGroupUsersCmd = &cobra.Command{
	Use:   "list-traffic-distribution-group-users",
	Short: "Lists traffic distribution group users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listTrafficDistributionGroupUsersCmd).Standalone()

	connect_listTrafficDistributionGroupUsersCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listTrafficDistributionGroupUsersCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listTrafficDistributionGroupUsersCmd.Flags().String("traffic-distribution-group-id", "", "The identifier of the traffic distribution group.")
	connect_listTrafficDistributionGroupUsersCmd.MarkFlagRequired("traffic-distribution-group-id")
	connectCmd.AddCommand(connect_listTrafficDistributionGroupUsersCmd)
}
