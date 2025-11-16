package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_submitFeedbackCmd = &cobra.Command{
	Use:   "submit-feedback",
	Short: "Enables you to provide feedback to Amazon Kendra to improve the performance of your index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_submitFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_submitFeedbackCmd).Standalone()

		kendra_submitFeedbackCmd.Flags().String("click-feedback-items", "", "Tells Amazon Kendra that a particular search result link was chosen by the user.")
		kendra_submitFeedbackCmd.Flags().String("index-id", "", "The identifier of the index that was queried.")
		kendra_submitFeedbackCmd.Flags().String("query-id", "", "The identifier of the specific query for which you are submitting feedback.")
		kendra_submitFeedbackCmd.Flags().String("relevance-feedback-items", "", "Provides Amazon Kendra with relevant or not relevant feedback for whether a particular item was relevant to the search.")
		kendra_submitFeedbackCmd.MarkFlagRequired("index-id")
		kendra_submitFeedbackCmd.MarkFlagRequired("query-id")
	})
	kendraCmd.AddCommand(kendra_submitFeedbackCmd)
}
