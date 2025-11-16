package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifySpotFleetRequestCmd = &cobra.Command{
	Use:   "modify-spot-fleet-request",
	Short: "Modifies the specified Spot Fleet request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifySpotFleetRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifySpotFleetRequestCmd).Standalone()

		ec2_modifySpotFleetRequestCmd.Flags().String("context", "", "Reserved.")
		ec2_modifySpotFleetRequestCmd.Flags().String("excess-capacity-termination-policy", "", "Indicates whether running instances should be terminated if the target capacity of the Spot Fleet request is decreased below the current size of the Spot Fleet.")
		ec2_modifySpotFleetRequestCmd.Flags().String("launch-template-configs", "", "The launch template and overrides.")
		ec2_modifySpotFleetRequestCmd.Flags().String("on-demand-target-capacity", "", "The number of On-Demand Instances in the fleet.")
		ec2_modifySpotFleetRequestCmd.Flags().String("spot-fleet-request-id", "", "The ID of the Spot Fleet request.")
		ec2_modifySpotFleetRequestCmd.Flags().String("target-capacity", "", "The size of the fleet.")
		ec2_modifySpotFleetRequestCmd.MarkFlagRequired("spot-fleet-request-id")
	})
	ec2Cmd.AddCommand(ec2_modifySpotFleetRequestCmd)
}
