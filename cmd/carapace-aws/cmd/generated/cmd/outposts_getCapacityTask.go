package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getCapacityTaskCmd = &cobra.Command{
	Use:   "get-capacity-task",
	Short: "Gets details of the specified capacity task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getCapacityTaskCmd).Standalone()

	outposts_getCapacityTaskCmd.Flags().String("capacity-task-id", "", "ID of the capacity task.")
	outposts_getCapacityTaskCmd.Flags().String("outpost-identifier", "", "ID or ARN of the Outpost associated with the specified capacity task.")
	outposts_getCapacityTaskCmd.MarkFlagRequired("capacity-task-id")
	outposts_getCapacityTaskCmd.MarkFlagRequired("outpost-identifier")
	outpostsCmd.AddCommand(outposts_getCapacityTaskCmd)
}
