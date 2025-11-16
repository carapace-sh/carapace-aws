package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTopicPermissionsCmd = &cobra.Command{
	Use:   "describe-topic-permissions",
	Short: "Describes the permissions of a topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTopicPermissionsCmd).Standalone()

	quicksight_describeTopicPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic that you want described.")
	quicksight_describeTopicPermissionsCmd.Flags().String("topic-id", "", "The ID of the topic that you want to describe.")
	quicksight_describeTopicPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeTopicPermissionsCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_describeTopicPermissionsCmd)
}
