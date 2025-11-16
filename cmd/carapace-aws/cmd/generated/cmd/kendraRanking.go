package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRankingCmd = &cobra.Command{
	Use:   "kendra-ranking",
	Short: "Amazon Kendra Intelligent Ranking uses Amazon Kendra semantic search capabilities to intelligently re-rank a search service's results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRankingCmd).Standalone()

	rootCmd.AddCommand(kendraRankingCmd)
}
