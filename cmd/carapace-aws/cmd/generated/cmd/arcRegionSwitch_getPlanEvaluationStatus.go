package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_getPlanEvaluationStatusCmd = &cobra.Command{
	Use:   "get-plan-evaluation-status",
	Short: "Retrieves the evaluation status of a Region switch plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_getPlanEvaluationStatusCmd).Standalone()

	arcRegionSwitch_getPlanEvaluationStatusCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	arcRegionSwitch_getPlanEvaluationStatusCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	arcRegionSwitch_getPlanEvaluationStatusCmd.Flags().String("plan-arn", "", "The Amazon Resource Name (ARN) of the Region switch plan to retrieve evaluation status for.")
	arcRegionSwitch_getPlanEvaluationStatusCmd.MarkFlagRequired("plan-arn")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_getPlanEvaluationStatusCmd)
}
