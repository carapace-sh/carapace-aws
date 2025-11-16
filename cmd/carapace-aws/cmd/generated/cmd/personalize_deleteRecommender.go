package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteRecommenderCmd = &cobra.Command{
	Use:   "delete-recommender",
	Short: "Deactivates and removes a recommender.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteRecommenderCmd).Standalone()

	personalize_deleteRecommenderCmd.Flags().String("recommender-arn", "", "The Amazon Resource Name (ARN) of the recommender to delete.")
	personalize_deleteRecommenderCmd.MarkFlagRequired("recommender-arn")
	personalizeCmd.AddCommand(personalize_deleteRecommenderCmd)
}
