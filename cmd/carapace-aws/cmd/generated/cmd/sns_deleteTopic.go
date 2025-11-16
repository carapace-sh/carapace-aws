package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_deleteTopicCmd = &cobra.Command{
	Use:   "delete-topic",
	Short: "Deletes a topic and all its subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_deleteTopicCmd).Standalone()

	sns_deleteTopicCmd.Flags().String("topic-arn", "", "The ARN of the topic you want to delete.")
	sns_deleteTopicCmd.MarkFlagRequired("topic-arn")
	snsCmd.AddCommand(sns_deleteTopicCmd)
}
