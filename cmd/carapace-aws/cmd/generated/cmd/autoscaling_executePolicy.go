package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_executePolicyCmd = &cobra.Command{
	Use:   "execute-policy",
	Short: "Executes the specified policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_executePolicyCmd).Standalone()

	autoscaling_executePolicyCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_executePolicyCmd.Flags().String("breach-threshold", "", "The breach threshold for the alarm.")
	autoscaling_executePolicyCmd.Flags().String("honor-cooldown", "", "Indicates whether Amazon EC2 Auto Scaling waits for the cooldown period to complete before executing the policy.")
	autoscaling_executePolicyCmd.Flags().String("metric-value", "", "The metric value to compare to `BreachThreshold`.")
	autoscaling_executePolicyCmd.Flags().String("policy-name", "", "The name or ARN of the policy.")
	autoscaling_executePolicyCmd.MarkFlagRequired("policy-name")
	autoscalingCmd.AddCommand(autoscaling_executePolicyCmd)
}
