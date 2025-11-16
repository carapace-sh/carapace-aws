package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_cancelPlanExecutionCmd = &cobra.Command{
	Use:   "cancel-plan-execution",
	Short: "Cancels an in-progress plan execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_cancelPlanExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_cancelPlanExecutionCmd).Standalone()

		arcRegionSwitch_cancelPlanExecutionCmd.Flags().String("comment", "", "A comment that you can enter about canceling a plan execution step.")
		arcRegionSwitch_cancelPlanExecutionCmd.Flags().String("execution-id", "", "The execution identifier of a plan execution.")
		arcRegionSwitch_cancelPlanExecutionCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the plan.")
		arcRegionSwitch_cancelPlanExecutionCmd.MarkFlagRequired("execution-id")
		arcRegionSwitch_cancelPlanExecutionCmd.MarkFlagRequired("plan-arn")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_cancelPlanExecutionCmd)
}
