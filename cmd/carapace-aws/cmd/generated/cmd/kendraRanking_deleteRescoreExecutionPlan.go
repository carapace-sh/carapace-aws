package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_deleteRescoreExecutionPlanCmd = &cobra.Command{
	Use:   "delete-rescore-execution-plan",
	Short: "Deletes a rescore execution plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_deleteRescoreExecutionPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendraRanking_deleteRescoreExecutionPlanCmd).Standalone()

		kendraRanking_deleteRescoreExecutionPlanCmd.Flags().String("id", "", "The identifier of the rescore execution plan that you want to delete.")
		kendraRanking_deleteRescoreExecutionPlanCmd.MarkFlagRequired("id")
	})
	kendraRankingCmd.AddCommand(kendraRanking_deleteRescoreExecutionPlanCmd)
}
