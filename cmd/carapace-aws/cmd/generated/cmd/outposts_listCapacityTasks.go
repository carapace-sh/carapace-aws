package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listCapacityTasksCmd = &cobra.Command{
	Use:   "list-capacity-tasks",
	Short: "Lists the capacity tasks for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listCapacityTasksCmd).Standalone()

	outposts_listCapacityTasksCmd.Flags().String("capacity-task-status-filter", "", "A list of statuses.")
	outposts_listCapacityTasksCmd.Flags().String("max-results", "", "")
	outposts_listCapacityTasksCmd.Flags().String("next-token", "", "")
	outposts_listCapacityTasksCmd.Flags().String("outpost-identifier-filter", "", "Filters the results by an Outpost ID or an Outpost ARN.")
	outpostsCmd.AddCommand(outposts_listCapacityTasksCmd)
}
