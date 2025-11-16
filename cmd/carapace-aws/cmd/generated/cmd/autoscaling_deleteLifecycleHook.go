package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deleteLifecycleHookCmd = &cobra.Command{
	Use:   "delete-lifecycle-hook",
	Short: "Deletes the specified lifecycle hook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deleteLifecycleHookCmd).Standalone()

	autoscaling_deleteLifecycleHookCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_deleteLifecycleHookCmd.Flags().String("lifecycle-hook-name", "", "The name of the lifecycle hook.")
	autoscaling_deleteLifecycleHookCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_deleteLifecycleHookCmd.MarkFlagRequired("lifecycle-hook-name")
	autoscalingCmd.AddCommand(autoscaling_deleteLifecycleHookCmd)
}
