package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_updatePlanExecutionStepCmd = &cobra.Command{
	Use:   "update-plan-execution-step",
	Short: "Updates a specific step in an in-progress plan execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_updatePlanExecutionStepCmd).Standalone()

	arcRegionSwitch_updatePlanExecutionStepCmd.Flags().String("action-to-take", "", "The updated action to take for the step.")
	arcRegionSwitch_updatePlanExecutionStepCmd.Flags().String("comment", "", "An optional comment about the plan execution.")
	arcRegionSwitch_updatePlanExecutionStepCmd.Flags().String("execution-id", "", "The unique identifier of the plan execution containing the step to update.")
	arcRegionSwitch_updatePlanExecutionStepCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the plan containing the execution step to update.")
	arcRegionSwitch_updatePlanExecutionStepCmd.Flags().String("step-name", "", "The name of the execution step to update.")
	arcRegionSwitch_updatePlanExecutionStepCmd.MarkFlagRequired("action-to-take")
	arcRegionSwitch_updatePlanExecutionStepCmd.MarkFlagRequired("comment")
	arcRegionSwitch_updatePlanExecutionStepCmd.MarkFlagRequired("execution-id")
	arcRegionSwitch_updatePlanExecutionStepCmd.MarkFlagRequired("plan-arn")
	arcRegionSwitch_updatePlanExecutionStepCmd.MarkFlagRequired("step-name")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_updatePlanExecutionStepCmd)
}
