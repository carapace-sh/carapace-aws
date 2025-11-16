package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeAutoScalingInstancesCmd = &cobra.Command{
	Use:   "describe-auto-scaling-instances",
	Short: "Gets information about the Auto Scaling instances in the account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeAutoScalingInstancesCmd).Standalone()

	autoscaling_describeAutoScalingInstancesCmd.Flags().String("instance-ids", "", "The IDs of the instances.")
	autoscaling_describeAutoScalingInstancesCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
	autoscaling_describeAutoScalingInstancesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	autoscalingCmd.AddCommand(autoscaling_describeAutoScalingInstancesCmd)
}
