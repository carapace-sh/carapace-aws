package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_putWarmPoolCmd = &cobra.Command{
	Use:   "put-warm-pool",
	Short: "Creates or updates a warm pool for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_putWarmPoolCmd).Standalone()

	autoscaling_putWarmPoolCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_putWarmPoolCmd.Flags().String("instance-reuse-policy", "", "Indicates whether instances in the Auto Scaling group can be returned to the warm pool on scale in.")
	autoscaling_putWarmPoolCmd.Flags().String("max-group-prepared-capacity", "", "Specifies the maximum number of instances that are allowed to be in the warm pool or in any state except `Terminated` for the Auto Scaling group.")
	autoscaling_putWarmPoolCmd.Flags().String("min-size", "", "Specifies the minimum number of instances to maintain in the warm pool.")
	autoscaling_putWarmPoolCmd.Flags().String("pool-state", "", "Sets the instance state to transition to after the lifecycle actions are complete.")
	autoscaling_putWarmPoolCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscalingCmd.AddCommand(autoscaling_putWarmPoolCmd)
}
