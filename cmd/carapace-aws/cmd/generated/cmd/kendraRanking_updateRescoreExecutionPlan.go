package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_updateRescoreExecutionPlanCmd = &cobra.Command{
	Use:   "update-rescore-execution-plan",
	Short: "Updates a rescore execution plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_updateRescoreExecutionPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendraRanking_updateRescoreExecutionPlanCmd).Standalone()

		kendraRanking_updateRescoreExecutionPlanCmd.Flags().String("capacity-units", "", "You can set additional capacity units to meet the needs of your rescore execution plan.")
		kendraRanking_updateRescoreExecutionPlanCmd.Flags().String("description", "", "A new description for the rescore execution plan.")
		kendraRanking_updateRescoreExecutionPlanCmd.Flags().String("id", "", "The identifier of the rescore execution plan that you want to update.")
		kendraRanking_updateRescoreExecutionPlanCmd.Flags().String("name", "", "A new name for the rescore execution plan.")
		kendraRanking_updateRescoreExecutionPlanCmd.MarkFlagRequired("id")
	})
	kendraRankingCmd.AddCommand(kendraRanking_updateRescoreExecutionPlanCmd)
}
