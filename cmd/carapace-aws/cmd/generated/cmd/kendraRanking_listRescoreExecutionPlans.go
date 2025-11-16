package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_listRescoreExecutionPlansCmd = &cobra.Command{
	Use:   "list-rescore-execution-plans",
	Short: "Lists your rescore execution plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_listRescoreExecutionPlansCmd).Standalone()

	kendraRanking_listRescoreExecutionPlansCmd.Flags().String("max-results", "", "The maximum number of rescore execution plans to return.")
	kendraRanking_listRescoreExecutionPlansCmd.Flags().String("next-token", "", "If the response is truncated, Amazon Kendra Intelligent Ranking returns a pagination token in the response.")
	kendraRankingCmd.AddCommand(kendraRanking_listRescoreExecutionPlansCmd)
}
