package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTopicCmd = &cobra.Command{
	Use:   "describe-topic",
	Short: "Describes a topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTopicCmd).Standalone()

	quicksight_describeTopicCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_describeTopicCmd.Flags().String("topic-id", "", "The ID of the topic that you want to describe.")
	quicksight_describeTopicCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeTopicCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_describeTopicCmd)
}
