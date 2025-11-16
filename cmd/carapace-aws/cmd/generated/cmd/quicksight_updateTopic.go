package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateTopicCmd = &cobra.Command{
	Use:   "update-topic",
	Short: "Updates a topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateTopicCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateTopicCmd).Standalone()

		quicksight_updateTopicCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic that you want to update.")
		quicksight_updateTopicCmd.Flags().String("custom-instructions", "", "Custom instructions for the topic.")
		quicksight_updateTopicCmd.Flags().String("topic", "", "The definition of the topic that you want to update.")
		quicksight_updateTopicCmd.Flags().String("topic-id", "", "The ID of the topic that you want to modify.")
		quicksight_updateTopicCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateTopicCmd.MarkFlagRequired("topic")
		quicksight_updateTopicCmd.MarkFlagRequired("topic-id")
	})
	quicksightCmd.AddCommand(quicksight_updateTopicCmd)
}
