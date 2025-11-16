package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getAccessLogSubscriptionCmd = &cobra.Command{
	Use:   "get-access-log-subscription",
	Short: "Retrieves information about the specified access log subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getAccessLogSubscriptionCmd).Standalone()

	vpcLattice_getAccessLogSubscriptionCmd.Flags().String("access-log-subscription-identifier", "", "The ID or ARN of the access log subscription.")
	vpcLattice_getAccessLogSubscriptionCmd.MarkFlagRequired("access-log-subscription-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getAccessLogSubscriptionCmd)
}
