package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeLifecycleHooksCmd = &cobra.Command{
	Use:   "describe-lifecycle-hooks",
	Short: "Gets information about the lifecycle hooks for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeLifecycleHooksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeLifecycleHooksCmd).Standalone()

		autoscaling_describeLifecycleHooksCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_describeLifecycleHooksCmd.Flags().String("lifecycle-hook-names", "", "The names of one or more lifecycle hooks.")
		autoscaling_describeLifecycleHooksCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_describeLifecycleHooksCmd)
}
