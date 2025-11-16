package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_listPlanExecutionEventsCmd = &cobra.Command{
	Use:   "list-plan-execution-events",
	Short: "Lists the events that occurred during a plan execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_listPlanExecutionEventsCmd).Standalone()

	arcRegionSwitch_listPlanExecutionEventsCmd.Flags().String("execution-id", "", "The execution identifier of a plan execution.")
	arcRegionSwitch_listPlanExecutionEventsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	arcRegionSwitch_listPlanExecutionEventsCmd.Flags().String("name", "", "The name of the plan execution event.")
	arcRegionSwitch_listPlanExecutionEventsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	arcRegionSwitch_listPlanExecutionEventsCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the plan.")
	arcRegionSwitch_listPlanExecutionEventsCmd.MarkFlagRequired("execution-id")
	arcRegionSwitch_listPlanExecutionEventsCmd.MarkFlagRequired("plan-arn")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_listPlanExecutionEventsCmd)
}
