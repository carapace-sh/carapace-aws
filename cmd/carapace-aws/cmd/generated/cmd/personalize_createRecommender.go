package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createRecommenderCmd = &cobra.Command{
	Use:   "create-recommender",
	Short: "Creates a recommender with the recipe (a Domain dataset group use case) you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createRecommenderCmd).Standalone()

	personalize_createRecommenderCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the destination domain dataset group for the recommender.")
	personalize_createRecommenderCmd.Flags().String("name", "", "The name of the recommender.")
	personalize_createRecommenderCmd.Flags().String("recipe-arn", "", "The Amazon Resource Name (ARN) of the recipe that the recommender will use.")
	personalize_createRecommenderCmd.Flags().String("recommender-config", "", "The configuration details of the recommender.")
	personalize_createRecommenderCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the recommender.")
	personalize_createRecommenderCmd.MarkFlagRequired("dataset-group-arn")
	personalize_createRecommenderCmd.MarkFlagRequired("name")
	personalize_createRecommenderCmd.MarkFlagRequired("recipe-arn")
	personalizeCmd.AddCommand(personalize_createRecommenderCmd)
}
