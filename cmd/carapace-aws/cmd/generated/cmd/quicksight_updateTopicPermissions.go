package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateTopicPermissionsCmd = &cobra.Command{
	Use:   "update-topic-permissions",
	Short: "Updates the permissions of a topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateTopicPermissionsCmd).Standalone()

	quicksight_updateTopicPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the topic that you want to update the permissions for.")
	quicksight_updateTopicPermissionsCmd.Flags().String("grant-permissions", "", "The resource permissions that you want to grant to the topic.")
	quicksight_updateTopicPermissionsCmd.Flags().String("revoke-permissions", "", "The resource permissions that you want to revoke from the topic.")
	quicksight_updateTopicPermissionsCmd.Flags().String("topic-id", "", "The ID of the topic that you want to modify.")
	quicksight_updateTopicPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateTopicPermissionsCmd.MarkFlagRequired("topic-id")
	quicksightCmd.AddCommand(quicksight_updateTopicPermissionsCmd)
}
