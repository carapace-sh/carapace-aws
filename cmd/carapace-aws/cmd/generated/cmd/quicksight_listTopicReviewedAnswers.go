package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listTopicReviewedAnswersCmd = &cobra.Command{
	Use:   "list-topic-reviewed-answers",
	Short: "Lists all reviewed answers for a Q Topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listTopicReviewedAnswersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listTopicReviewedAnswersCmd).Standalone()

		quicksight_listTopicReviewedAnswersCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that containd the reviewed answers that you want listed.")
		quicksight_listTopicReviewedAnswersCmd.Flags().String("topic-id", "", "The ID for the topic that contains the reviewed answer that you want to list.")
		quicksight_listTopicReviewedAnswersCmd.MarkFlagRequired("aws-account-id")
		quicksight_listTopicReviewedAnswersCmd.MarkFlagRequired("topic-id")
	})
	quicksightCmd.AddCommand(quicksight_listTopicReviewedAnswersCmd)
}
