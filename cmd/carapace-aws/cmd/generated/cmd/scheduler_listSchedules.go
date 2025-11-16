package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_listSchedulesCmd = &cobra.Command{
	Use:   "list-schedules",
	Short: "Returns a paginated list of your EventBridge Scheduler schedules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_listSchedulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_listSchedulesCmd).Standalone()

		scheduler_listSchedulesCmd.Flags().String("group-name", "", "If specified, only lists the schedules whose associated schedule group matches the given filter.")
		scheduler_listSchedulesCmd.Flags().String("max-results", "", "If specified, limits the number of results returned by this operation.")
		scheduler_listSchedulesCmd.Flags().String("name-prefix", "", "Schedule name prefix to return the filtered list of resources.")
		scheduler_listSchedulesCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
		scheduler_listSchedulesCmd.Flags().String("state", "", "If specified, only lists the schedules whose current state matches the given filter.")
	})
	schedulerCmd.AddCommand(scheduler_listSchedulesCmd)
}
