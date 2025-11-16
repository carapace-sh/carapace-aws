package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_stopRecommenderCmd = &cobra.Command{
	Use:   "stop-recommender",
	Short: "Stops a recommender that is ACTIVE.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_stopRecommenderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_stopRecommenderCmd).Standalone()

		personalize_stopRecommenderCmd.Flags().String("recommender-arn", "", "The Amazon Resource Name (ARN) of the recommender to stop.")
		personalize_stopRecommenderCmd.MarkFlagRequired("recommender-arn")
	})
	personalizeCmd.AddCommand(personalize_stopRecommenderCmd)
}
