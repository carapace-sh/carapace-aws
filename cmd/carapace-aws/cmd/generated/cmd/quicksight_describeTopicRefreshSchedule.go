package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTopicRefreshScheduleCmd = &cobra.Command{
	Use:   "describe-topic-refresh-schedule",
	Short: "Deletes a topic refresh schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTopicRefreshScheduleCmd).Standalone()

	quicksight_describeTopicRefreshScheduleCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_describeTopicRefreshScheduleCmd.Flags().String("dataset-id", "", "The ID of the dataset.")
	quicksight_describeTopicRefreshScheduleCmd.Flags().String("topic-id", "", "The ID of the topic that contains the refresh schedule that you want to describe.")
	quicksight_describeTopicRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeTopicRefreshScheduleCmd.MarkFlagRequired("dataset-id")
	quicksight_describeTopicRefreshScheduleCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_describeTopicRefreshScheduleCmd)
}
