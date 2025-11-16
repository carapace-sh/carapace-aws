package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_getScheduleCmd = &cobra.Command{
	Use:   "get-schedule",
	Short: "Retrieves the specified schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_getScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_getScheduleCmd).Standalone()

		scheduler_getScheduleCmd.Flags().String("group-name", "", "The name of the schedule group associated with this schedule.")
		scheduler_getScheduleCmd.Flags().String("name", "", "The name of the schedule to retrieve.")
		scheduler_getScheduleCmd.MarkFlagRequired("name")
	})
	schedulerCmd.AddCommand(scheduler_getScheduleCmd)
}
