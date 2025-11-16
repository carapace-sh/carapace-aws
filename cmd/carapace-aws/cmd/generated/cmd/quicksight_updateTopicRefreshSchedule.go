package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateTopicRefreshScheduleCmd = &cobra.Command{
	Use:   "update-topic-refresh-schedule",
	Short: "Updates a topic refresh schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateTopicRefreshScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateTopicRefreshScheduleCmd).Standalone()

		quicksight_updateTopicRefreshScheduleCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic whose refresh schedule you want to update.")
		quicksight_updateTopicRefreshScheduleCmd.Flags().String("dataset-id", "", "The ID of the dataset.")
		quicksight_updateTopicRefreshScheduleCmd.Flags().String("refresh-schedule", "", "The definition of a refresh schedule.")
		quicksight_updateTopicRefreshScheduleCmd.Flags().String("topic-id", "", "The ID of the topic that you want to modify.")
		quicksight_updateTopicRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateTopicRefreshScheduleCmd.MarkFlagRequired("dataset-id")
		quicksight_updateTopicRefreshScheduleCmd.MarkFlagRequired("refresh-schedule")
		quicksight_updateTopicRefreshScheduleCmd.MarkFlagRequired("topic-id")
	})
	quicksightCmd.AddCommand(quicksight_updateTopicRefreshScheduleCmd)
}
