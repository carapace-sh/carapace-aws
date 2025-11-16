package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listAccessLogSubscriptionsCmd = &cobra.Command{
	Use:   "list-access-log-subscriptions",
	Short: "Lists the access log subscriptions for the specified service network or service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listAccessLogSubscriptionsCmd).Standalone()

	vpcLattice_listAccessLogSubscriptionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	vpcLattice_listAccessLogSubscriptionsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLattice_listAccessLogSubscriptionsCmd.Flags().String("resource-identifier", "", "The ID or ARN of the service network or service.")
	vpcLattice_listAccessLogSubscriptionsCmd.MarkFlagRequired("resource-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_listAccessLogSubscriptionsCmd)
}
