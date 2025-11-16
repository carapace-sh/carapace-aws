package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateKeyRegistrationCmd = &cobra.Command{
	Use:   "update-key-registration",
	Short: "Updates a customer managed key in a Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateKeyRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateKeyRegistrationCmd).Standalone()

		quicksight_updateKeyRegistrationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the customer managed key registration that you want to update.")
		quicksight_updateKeyRegistrationCmd.Flags().String("key-registration", "", "A list of `RegisteredCustomerManagedKey` objects to be updated to the Quick Sight account.")
		quicksight_updateKeyRegistrationCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateKeyRegistrationCmd.MarkFlagRequired("key-registration")
	})
	quicksightCmd.AddCommand(quicksight_updateKeyRegistrationCmd)
}
