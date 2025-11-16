package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_getPlanExecutionCmd = &cobra.Command{
	Use:   "get-plan-execution",
	Short: "Retrieves detailed information about a specific plan execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_getPlanExecutionCmd).Standalone()

	arcRegionSwitch_getPlanExecutionCmd.Flags().String("execution-id", "", "The execution identifier of a plan execution.")
	arcRegionSwitch_getPlanExecutionCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	arcRegionSwitch_getPlanExecutionCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	arcRegionSwitch_getPlanExecutionCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the plan with the execution to retrieve.")
	arcRegionSwitch_getPlanExecutionCmd.MarkFlagRequired("execution-id")
	arcRegionSwitch_getPlanExecutionCmd.MarkFlagRequired("plan-arn")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_getPlanExecutionCmd)
}
