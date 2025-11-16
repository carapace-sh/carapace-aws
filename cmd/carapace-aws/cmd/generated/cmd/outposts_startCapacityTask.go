package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_startCapacityTaskCmd = &cobra.Command{
	Use:   "start-capacity-task",
	Short: "Starts the specified capacity task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_startCapacityTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_startCapacityTaskCmd).Standalone()

		outposts_startCapacityTaskCmd.Flags().String("asset-id", "", "The ID of the Outpost asset.")
		outposts_startCapacityTaskCmd.Flags().String("dry-run", "", "You can request a dry run to determine if the instance type and instance size changes is above or below available instance capacity.")
		outposts_startCapacityTaskCmd.Flags().String("instance-pools", "", "The instance pools specified in the capacity task.")
		outposts_startCapacityTaskCmd.Flags().String("instances-to-exclude", "", "List of user-specified running instances that must not be stopped in order to free up the capacity needed to run the capacity task.")
		outposts_startCapacityTaskCmd.Flags().String("order-id", "", "The ID of the Amazon Web Services Outposts order associated with the specified capacity task.")
		outposts_startCapacityTaskCmd.Flags().String("outpost-identifier", "", "The ID or ARN of the Outposts associated with the specified capacity task.")
		outposts_startCapacityTaskCmd.Flags().String("task-action-on-blocking-instances", "", "Specify one of the following options in case an instance is blocking the capacity task from running.")
		outposts_startCapacityTaskCmd.MarkFlagRequired("instance-pools")
		outposts_startCapacityTaskCmd.MarkFlagRequired("outpost-identifier")
	})
	outpostsCmd.AddCommand(outposts_startCapacityTaskCmd)
}
