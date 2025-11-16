package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteTopicRefreshScheduleCmd = &cobra.Command{
	Use:   "delete-topic-refresh-schedule",
	Short: "Deletes a topic refresh schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteTopicRefreshScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteTopicRefreshScheduleCmd).Standalone()

		quicksight_deleteTopicRefreshScheduleCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_deleteTopicRefreshScheduleCmd.Flags().String("dataset-id", "", "The ID of the dataset.")
		quicksight_deleteTopicRefreshScheduleCmd.Flags().String("topic-id", "", "The ID of the topic that you want to modify.")
		quicksight_deleteTopicRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteTopicRefreshScheduleCmd.MarkFlagRequired("dataset-id")
		quicksight_deleteTopicRefreshScheduleCmd.MarkFlagRequired("topic-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteTopicRefreshScheduleCmd)
}
