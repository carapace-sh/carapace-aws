package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_listScheduleGroupsCmd = &cobra.Command{
	Use:   "list-schedule-groups",
	Short: "Returns a paginated list of your schedule groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_listScheduleGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_listScheduleGroupsCmd).Standalone()

		scheduler_listScheduleGroupsCmd.Flags().String("max-results", "", "If specified, limits the number of results returned by this operation.")
		scheduler_listScheduleGroupsCmd.Flags().String("name-prefix", "", "The name prefix that you can use to return a filtered list of your schedule groups.")
		scheduler_listScheduleGroupsCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
	})
	schedulerCmd.AddCommand(scheduler_listScheduleGroupsCmd)
}
