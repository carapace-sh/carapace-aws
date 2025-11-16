package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_getScheduleGroupCmd = &cobra.Command{
	Use:   "get-schedule-group",
	Short: "Retrieves the specified schedule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_getScheduleGroupCmd).Standalone()

	scheduler_getScheduleGroupCmd.Flags().String("name", "", "The name of the schedule group to retrieve.")
	scheduler_getScheduleGroupCmd.MarkFlagRequired("name")
	schedulerCmd.AddCommand(scheduler_getScheduleGroupCmd)
}
