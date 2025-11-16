package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_approvePlanExecutionStepCmd = &cobra.Command{
	Use:   "approve-plan-execution-step",
	Short: "Approves a step in a plan execution that requires manual approval.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_approvePlanExecutionStepCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_approvePlanExecutionStepCmd).Standalone()

		arcRegionSwitch_approvePlanExecutionStepCmd.Flags().String("approval", "", "The status of approval for a plan execution step.")
		arcRegionSwitch_approvePlanExecutionStepCmd.Flags().String("comment", "", "A comment that you can enter about a plan execution.")
		arcRegionSwitch_approvePlanExecutionStepCmd.Flags().String("execution-id", "", "The execution identifier of a plan execution.")
		arcRegionSwitch_approvePlanExecutionStepCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the plan.")
		arcRegionSwitch_approvePlanExecutionStepCmd.Flags().String("step-name", "", "The name of a step in a plan execution.")
		arcRegionSwitch_approvePlanExecutionStepCmd.MarkFlagRequired("approval")
		arcRegionSwitch_approvePlanExecutionStepCmd.MarkFlagRequired("execution-id")
		arcRegionSwitch_approvePlanExecutionStepCmd.MarkFlagRequired("plan-arn")
		arcRegionSwitch_approvePlanExecutionStepCmd.MarkFlagRequired("step-name")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_approvePlanExecutionStepCmd)
}
