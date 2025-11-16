package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_batchDeleteTopicReviewedAnswerCmd = &cobra.Command{
	Use:   "batch-delete-topic-reviewed-answer",
	Short: "Deletes reviewed answers for Q Topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_batchDeleteTopicReviewedAnswerCmd).Standalone()

	quicksight_batchDeleteTopicReviewedAnswerCmd.Flags().String("answer-ids", "", "The Answer IDs of the Answers to be deleted.")
	quicksight_batchDeleteTopicReviewedAnswerCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that you want to delete a reviewed answers in.")
	quicksight_batchDeleteTopicReviewedAnswerCmd.Flags().String("topic-id", "", "The ID for the topic reviewed answer that you want to delete.")
	quicksight_batchDeleteTopicReviewedAnswerCmd.MarkFlagRequired("aws-account-id")
	quicksight_batchDeleteTopicReviewedAnswerCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_batchDeleteTopicReviewedAnswerCmd)
}
