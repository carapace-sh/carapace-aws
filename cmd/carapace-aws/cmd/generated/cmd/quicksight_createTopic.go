package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createTopicCmd = &cobra.Command{
	Use:   "create-topic",
	Short: "Creates a new Q topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createTopicCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createTopicCmd).Standalone()

		quicksight_createTopicCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that you want to create a topic in.")
		quicksight_createTopicCmd.Flags().String("custom-instructions", "", "Custom instructions for the topic.")
		quicksight_createTopicCmd.Flags().String("folder-arns", "", "The Folder ARN of the folder that you want the topic to reside in.")
		quicksight_createTopicCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags that are assigned to the dataset.")
		quicksight_createTopicCmd.Flags().String("topic", "", "The definition of a topic to create.")
		quicksight_createTopicCmd.Flags().String("topic-id", "", "The ID for the topic that you want to create.")
		quicksight_createTopicCmd.MarkFlagRequired("aws-account-id")
		quicksight_createTopicCmd.MarkFlagRequired("topic")
		quicksight_createTopicCmd.MarkFlagRequired("topic-id")
	})
	quicksightCmd.AddCommand(quicksight_createTopicCmd)
}
