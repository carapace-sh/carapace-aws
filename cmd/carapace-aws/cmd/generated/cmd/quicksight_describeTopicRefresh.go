package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTopicRefreshCmd = &cobra.Command{
	Use:   "describe-topic-refresh",
	Short: "Describes the status of a topic refresh.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTopicRefreshCmd).Standalone()

	quicksight_describeTopicRefreshCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic whose refresh you want to describe.")
	quicksight_describeTopicRefreshCmd.Flags().String("refresh-id", "", "The ID of the refresh, which is performed when the topic is created or updated.")
	quicksight_describeTopicRefreshCmd.Flags().String("topic-id", "", "The ID of the topic that you want to describe.")
	quicksight_describeTopicRefreshCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeTopicRefreshCmd.MarkFlagRequired("refresh-id")
	quicksight_describeTopicRefreshCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_describeTopicRefreshCmd)
}
