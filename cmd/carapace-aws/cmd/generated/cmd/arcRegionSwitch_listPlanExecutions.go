package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_listPlanExecutionsCmd = &cobra.Command{
	Use:   "list-plan-executions",
	Short: "Lists the executions of a Region switch plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_listPlanExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_listPlanExecutionsCmd).Standalone()

		arcRegionSwitch_listPlanExecutionsCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		arcRegionSwitch_listPlanExecutionsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		arcRegionSwitch_listPlanExecutionsCmd.Flags().String("plan-arn", "", "The ARN for the plan.")
		arcRegionSwitch_listPlanExecutionsCmd.Flags().String("state", "", "The state of the plan execution.")
		arcRegionSwitch_listPlanExecutionsCmd.MarkFlagRequired("plan-arn")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_listPlanExecutionsCmd)
}
