package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_describeRescoreExecutionPlanCmd = &cobra.Command{
	Use:   "describe-rescore-execution-plan",
	Short: "Gets information about a rescore execution plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_describeRescoreExecutionPlanCmd).Standalone()

	kendraRanking_describeRescoreExecutionPlanCmd.Flags().String("id", "", "The identifier of the rescore execution plan that you want to get information on.")
	kendraRanking_describeRescoreExecutionPlanCmd.MarkFlagRequired("id")
	kendraRankingCmd.AddCommand(kendraRanking_describeRescoreExecutionPlanCmd)
}
