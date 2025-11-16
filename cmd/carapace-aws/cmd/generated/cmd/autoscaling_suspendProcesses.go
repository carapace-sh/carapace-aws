package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_suspendProcessesCmd = &cobra.Command{
	Use:   "suspend-processes",
	Short: "Suspends the specified auto scaling processes, or all processes, for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_suspendProcessesCmd).Standalone()

	autoscaling_suspendProcessesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_suspendProcessesCmd.Flags().String("scaling-processes", "", "One or more of the following processes:")
	autoscaling_suspendProcessesCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscalingCmd.AddCommand(autoscaling_suspendProcessesCmd)
}
