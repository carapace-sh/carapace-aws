package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_startRecommenderCmd = &cobra.Command{
	Use:   "start-recommender",
	Short: "Starts a recommender that is INACTIVE.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_startRecommenderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_startRecommenderCmd).Standalone()

		personalize_startRecommenderCmd.Flags().String("recommender-arn", "", "The Amazon Resource Name (ARN) of the recommender to start.")
		personalize_startRecommenderCmd.MarkFlagRequired("recommender-arn")
	})
	personalizeCmd.AddCommand(personalize_startRecommenderCmd)
}
