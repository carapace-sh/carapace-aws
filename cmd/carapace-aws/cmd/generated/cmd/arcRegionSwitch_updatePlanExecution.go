package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_updatePlanExecutionCmd = &cobra.Command{
	Use:   "update-plan-execution",
	Short: "Updates an in-progress plan execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_updatePlanExecutionCmd).Standalone()

	arcRegionSwitch_updatePlanExecutionCmd.Flags().String("action", "", "The action specified for a plan execution, for example, Switch to Graceful or Pause.")
	arcRegionSwitch_updatePlanExecutionCmd.Flags().String("comment", "", "An optional comment about the plan execution.")
	arcRegionSwitch_updatePlanExecutionCmd.Flags().String("execution-id", "", "The execution identifier of a plan execution.")
	arcRegionSwitch_updatePlanExecutionCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the plan with the execution to update.")
	arcRegionSwitch_updatePlanExecutionCmd.MarkFlagRequired("action")
	arcRegionSwitch_updatePlanExecutionCmd.MarkFlagRequired("execution-id")
	arcRegionSwitch_updatePlanExecutionCmd.MarkFlagRequired("plan-arn")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_updatePlanExecutionCmd)
}
