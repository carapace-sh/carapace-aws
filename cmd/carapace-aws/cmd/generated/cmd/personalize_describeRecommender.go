package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeRecommenderCmd = &cobra.Command{
	Use:   "describe-recommender",
	Short: "Describes the given recommender, including its status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeRecommenderCmd).Standalone()

	personalize_describeRecommenderCmd.Flags().String("recommender-arn", "", "The Amazon Resource Name (ARN) of the recommender to describe.")
	personalize_describeRecommenderCmd.MarkFlagRequired("recommender-arn")
	personalizeCmd.AddCommand(personalize_describeRecommenderCmd)
}
