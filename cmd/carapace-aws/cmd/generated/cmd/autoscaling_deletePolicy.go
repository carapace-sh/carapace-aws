package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Deletes the specified scaling policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deletePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_deletePolicyCmd).Standalone()

		autoscaling_deletePolicyCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_deletePolicyCmd.Flags().String("policy-name", "", "The name or Amazon Resource Name (ARN) of the policy.")
		autoscaling_deletePolicyCmd.MarkFlagRequired("policy-name")
	})
	autoscalingCmd.AddCommand(autoscaling_deletePolicyCmd)
}
