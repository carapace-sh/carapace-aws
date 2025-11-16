package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_setInstanceHealthCmd = &cobra.Command{
	Use:   "set-instance-health",
	Short: "Sets the health status of the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_setInstanceHealthCmd).Standalone()

	autoscaling_setInstanceHealthCmd.Flags().String("health-status", "", "The health status of the instance.")
	autoscaling_setInstanceHealthCmd.Flags().String("instance-id", "", "The ID of the instance.")
	autoscaling_setInstanceHealthCmd.Flags().String("should-respect-grace-period", "", "If the Auto Scaling group of the specified instance has a `HealthCheckGracePeriod` specified for the group, by default, this call respects the grace period.")
	autoscaling_setInstanceHealthCmd.MarkFlagRequired("health-status")
	autoscaling_setInstanceHealthCmd.MarkFlagRequired("instance-id")
	autoscalingCmd.AddCommand(autoscaling_setInstanceHealthCmd)
}
