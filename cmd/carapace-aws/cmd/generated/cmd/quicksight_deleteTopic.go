package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteTopicCmd = &cobra.Command{
	Use:   "delete-topic",
	Short: "Deletes a topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteTopicCmd).Standalone()

	quicksight_deleteTopicCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic that you want to delete.")
	quicksight_deleteTopicCmd.Flags().String("topic-id", "", "The ID of the topic that you want to delete.")
	quicksight_deleteTopicCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteTopicCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_deleteTopicCmd)
}
