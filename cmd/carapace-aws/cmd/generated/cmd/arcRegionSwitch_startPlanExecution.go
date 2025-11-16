package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_startPlanExecutionCmd = &cobra.Command{
	Use:   "start-plan-execution",
	Short: "Starts the execution of a Region switch plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_startPlanExecutionCmd).Standalone()

	arcRegionSwitch_startPlanExecutionCmd.Flags().String("action", "", "The action to perform.")
	arcRegionSwitch_startPlanExecutionCmd.Flags().String("comment", "", "An optional comment explaining why the plan execution is being started.")
	arcRegionSwitch_startPlanExecutionCmd.Flags().String("latest-version", "", "A boolean value indicating whether to use the latest version of the plan.")
	arcRegionSwitch_startPlanExecutionCmd.Flags().String("mode", "", "The plan execution mode.")
	arcRegionSwitch_startPlanExecutionCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the plan to execute.")
	arcRegionSwitch_startPlanExecutionCmd.Flags().String("target-region", "", "The Amazon Web Services Region to target with this execution.")
	arcRegionSwitch_startPlanExecutionCmd.MarkFlagRequired("action")
	arcRegionSwitch_startPlanExecutionCmd.MarkFlagRequired("plan-arn")
	arcRegionSwitch_startPlanExecutionCmd.MarkFlagRequired("target-region")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_startPlanExecutionCmd)
}
