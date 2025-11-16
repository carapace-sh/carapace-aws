package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_rescoreCmd = &cobra.Command{
	Use:   "rescore",
	Short: "Rescores or re-ranks search results from a search service such as OpenSearch (self managed).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_rescoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendraRanking_rescoreCmd).Standalone()

		kendraRanking_rescoreCmd.Flags().String("documents", "", "The list of documents for Amazon Kendra Intelligent Ranking to rescore or rank on.")
		kendraRanking_rescoreCmd.Flags().String("rescore-execution-plan-id", "", "The identifier of the rescore execution plan.")
		kendraRanking_rescoreCmd.Flags().String("search-query", "", "The input query from the search service.")
		kendraRanking_rescoreCmd.MarkFlagRequired("documents")
		kendraRanking_rescoreCmd.MarkFlagRequired("rescore-execution-plan-id")
		kendraRanking_rescoreCmd.MarkFlagRequired("search-query")
	})
	kendraRankingCmd.AddCommand(kendraRanking_rescoreCmd)
}
