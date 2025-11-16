package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listTopicRefreshSchedulesCmd = &cobra.Command{
	Use:   "list-topic-refresh-schedules",
	Short: "Lists all of the refresh schedules for a topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listTopicRefreshSchedulesCmd).Standalone()

	quicksight_listTopicRefreshSchedulesCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic whose refresh schedule you want described.")
	quicksight_listTopicRefreshSchedulesCmd.Flags().String("topic-id", "", "The ID for the topic that you want to describe.")
	quicksight_listTopicRefreshSchedulesCmd.MarkFlagRequired("aws-account-id")
	quicksight_listTopicRefreshSchedulesCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_listTopicRefreshSchedulesCmd)
}
