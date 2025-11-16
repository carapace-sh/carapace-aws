package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteAccessLogSubscriptionCmd = &cobra.Command{
	Use:   "delete-access-log-subscription",
	Short: "Deletes the specified access log subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteAccessLogSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteAccessLogSubscriptionCmd).Standalone()

		vpcLattice_deleteAccessLogSubscriptionCmd.Flags().String("access-log-subscription-identifier", "", "The ID or ARN of the access log subscription.")
		vpcLattice_deleteAccessLogSubscriptionCmd.MarkFlagRequired("access-log-subscription-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteAccessLogSubscriptionCmd)
}
