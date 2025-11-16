package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteAccountSubscriptionCmd = &cobra.Command{
	Use:   "delete-account-subscription",
	Short: "Deleting your Quick Sight account subscription has permanent, irreversible consequences across all Amazon Web Services regions:\n\n- Global deletion – Running this operation from any single region will delete your Quick Sight account and all data in every Amazon Web Services region where you have Quick Sight resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteAccountSubscriptionCmd).Standalone()

	quicksight_deleteAccountSubscriptionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID of the account that you want to delete.")
	quicksight_deleteAccountSubscriptionCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_deleteAccountSubscriptionCmd)
}
