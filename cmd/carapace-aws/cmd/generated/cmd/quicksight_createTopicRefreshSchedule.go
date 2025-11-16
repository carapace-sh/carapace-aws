package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createTopicRefreshScheduleCmd = &cobra.Command{
	Use:   "create-topic-refresh-schedule",
	Short: "Creates a topic refresh schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createTopicRefreshScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createTopicRefreshScheduleCmd).Standalone()

		quicksight_createTopicRefreshScheduleCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic you're creating a refresh schedule for.")
		quicksight_createTopicRefreshScheduleCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset.")
		quicksight_createTopicRefreshScheduleCmd.Flags().String("dataset-name", "", "The name of the dataset.")
		quicksight_createTopicRefreshScheduleCmd.Flags().String("refresh-schedule", "", "The definition of a refresh schedule.")
		quicksight_createTopicRefreshScheduleCmd.Flags().String("topic-id", "", "The ID of the topic that you want to modify.")
		quicksight_createTopicRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
		quicksight_createTopicRefreshScheduleCmd.MarkFlagRequired("dataset-arn")
		quicksight_createTopicRefreshScheduleCmd.MarkFlagRequired("refresh-schedule")
		quicksight_createTopicRefreshScheduleCmd.MarkFlagRequired("topic-id")
	})
	quicksightCmd.AddCommand(quicksight_createTopicRefreshScheduleCmd)
}
