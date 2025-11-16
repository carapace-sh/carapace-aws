package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeAutoScalingGroupsCmd = &cobra.Command{
	Use:   "describe-auto-scaling-groups",
	Short: "Gets information about the Auto Scaling groups in the account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeAutoScalingGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeAutoScalingGroupsCmd).Standalone()

		autoscaling_describeAutoScalingGroupsCmd.Flags().String("auto-scaling-group-names", "", "The names of the Auto Scaling groups.")
		autoscaling_describeAutoScalingGroupsCmd.Flags().String("filters", "", "One or more filters to limit the results based on specific tags.")
		autoscaling_describeAutoScalingGroupsCmd.Flags().String("include-instances", "", "Specifies whether to include information about Amazon EC2 instances in the response.")
		autoscaling_describeAutoScalingGroupsCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
		autoscaling_describeAutoScalingGroupsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	autoscalingCmd.AddCommand(autoscaling_describeAutoScalingGroupsCmd)
}
