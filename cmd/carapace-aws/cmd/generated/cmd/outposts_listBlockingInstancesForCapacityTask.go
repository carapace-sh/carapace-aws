package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listBlockingInstancesForCapacityTaskCmd = &cobra.Command{
	Use:   "list-blocking-instances-for-capacity-task",
	Short: "A list of Amazon EC2 instances running on the Outpost and belonging to the account that initiated the capacity task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listBlockingInstancesForCapacityTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_listBlockingInstancesForCapacityTaskCmd).Standalone()

		outposts_listBlockingInstancesForCapacityTaskCmd.Flags().String("capacity-task-id", "", "The ID of the capacity task.")
		outposts_listBlockingInstancesForCapacityTaskCmd.Flags().String("max-results", "", "")
		outposts_listBlockingInstancesForCapacityTaskCmd.Flags().String("next-token", "", "")
		outposts_listBlockingInstancesForCapacityTaskCmd.Flags().String("outpost-identifier", "", "The ID or ARN of the Outpost associated with the specified capacity task.")
		outposts_listBlockingInstancesForCapacityTaskCmd.MarkFlagRequired("capacity-task-id")
		outposts_listBlockingInstancesForCapacityTaskCmd.MarkFlagRequired("outpost-identifier")
	})
	outpostsCmd.AddCommand(outposts_listBlockingInstancesForCapacityTaskCmd)
}
