package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_updateRecommenderCmd = &cobra.Command{
	Use:   "update-recommender",
	Short: "Updates the recommender to modify the recommender configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_updateRecommenderCmd).Standalone()

	personalize_updateRecommenderCmd.Flags().String("recommender-arn", "", "The Amazon Resource Name (ARN) of the recommender to modify.")
	personalize_updateRecommenderCmd.Flags().String("recommender-config", "", "The configuration details of the recommender.")
	personalize_updateRecommenderCmd.MarkFlagRequired("recommender-arn")
	personalize_updateRecommenderCmd.MarkFlagRequired("recommender-config")
	personalizeCmd.AddCommand(personalize_updateRecommenderCmd)
}
