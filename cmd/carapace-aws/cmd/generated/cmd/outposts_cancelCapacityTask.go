package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_cancelCapacityTaskCmd = &cobra.Command{
	Use:   "cancel-capacity-task",
	Short: "Cancels the capacity task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_cancelCapacityTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_cancelCapacityTaskCmd).Standalone()

		outposts_cancelCapacityTaskCmd.Flags().String("capacity-task-id", "", "ID of the capacity task that you want to cancel.")
		outposts_cancelCapacityTaskCmd.Flags().String("outpost-identifier", "", "ID or ARN of the Outpost associated with the capacity task that you want to cancel.")
		outposts_cancelCapacityTaskCmd.MarkFlagRequired("capacity-task-id")
		outposts_cancelCapacityTaskCmd.MarkFlagRequired("outpost-identifier")
	})
	outpostsCmd.AddCommand(outposts_cancelCapacityTaskCmd)
}
