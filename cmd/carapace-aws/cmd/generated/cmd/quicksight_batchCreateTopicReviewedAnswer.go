package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_batchCreateTopicReviewedAnswerCmd = &cobra.Command{
	Use:   "batch-create-topic-reviewed-answer",
	Short: "Creates new reviewed answers for a Q Topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_batchCreateTopicReviewedAnswerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_batchCreateTopicReviewedAnswerCmd).Standalone()

		quicksight_batchCreateTopicReviewedAnswerCmd.Flags().String("answers", "", "The definition of the Answers to be created.")
		quicksight_batchCreateTopicReviewedAnswerCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that you want to create a reviewed answer in.")
		quicksight_batchCreateTopicReviewedAnswerCmd.Flags().String("topic-id", "", "The ID for the topic reviewed answer that you want to create.")
		quicksight_batchCreateTopicReviewedAnswerCmd.MarkFlagRequired("answers")
		quicksight_batchCreateTopicReviewedAnswerCmd.MarkFlagRequired("aws-account-id")
		quicksight_batchCreateTopicReviewedAnswerCmd.MarkFlagRequired("topic-id")
	})
	quicksightCmd.AddCommand(quicksight_batchCreateTopicReviewedAnswerCmd)
}
