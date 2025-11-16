package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateAccessLogSubscriptionCmd = &cobra.Command{
	Use:   "update-access-log-subscription",
	Short: "Updates the specified access log subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateAccessLogSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_updateAccessLogSubscriptionCmd).Standalone()

		vpcLattice_updateAccessLogSubscriptionCmd.Flags().String("access-log-subscription-identifier", "", "The ID or ARN of the access log subscription.")
		vpcLattice_updateAccessLogSubscriptionCmd.Flags().String("destination-arn", "", "The Amazon Resource Name (ARN) of the access log destination.")
		vpcLattice_updateAccessLogSubscriptionCmd.MarkFlagRequired("access-log-subscription-identifier")
		vpcLattice_updateAccessLogSubscriptionCmd.MarkFlagRequired("destination-arn")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_updateAccessLogSubscriptionCmd)
}
