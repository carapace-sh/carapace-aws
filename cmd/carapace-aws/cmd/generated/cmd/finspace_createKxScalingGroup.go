package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxScalingGroupCmd = &cobra.Command{
	Use:   "create-kx-scaling-group",
	Short: "Creates a new scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxScalingGroupCmd).Standalone()

	finspace_createKxScalingGroupCmd.Flags().String("availability-zone-id", "", "The identifier of the availability zones.")
	finspace_createKxScalingGroupCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_createKxScalingGroupCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, where you want to create the scaling group.")
	finspace_createKxScalingGroupCmd.Flags().String("host-type", "", "The memory and CPU capabilities of the scaling group host on which FinSpace Managed kdb clusters will be placed.")
	finspace_createKxScalingGroupCmd.Flags().String("scaling-group-name", "", "A unique identifier for the kdb scaling group.")
	finspace_createKxScalingGroupCmd.Flags().String("tags", "", "A list of key-value pairs to label the scaling group.")
	finspace_createKxScalingGroupCmd.MarkFlagRequired("availability-zone-id")
	finspace_createKxScalingGroupCmd.MarkFlagRequired("client-token")
	finspace_createKxScalingGroupCmd.MarkFlagRequired("environment-id")
	finspace_createKxScalingGroupCmd.MarkFlagRequired("host-type")
	finspace_createKxScalingGroupCmd.MarkFlagRequired("scaling-group-name")
	finspaceCmd.AddCommand(finspace_createKxScalingGroupCmd)
}
