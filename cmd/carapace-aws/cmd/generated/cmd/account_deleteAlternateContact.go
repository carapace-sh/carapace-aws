package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_deleteAlternateContactCmd = &cobra.Command{
	Use:   "delete-alternate-contact",
	Short: "Deletes the specified alternate contact from an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_deleteAlternateContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(account_deleteAlternateContactCmd).Standalone()

		account_deleteAlternateContactCmd.Flags().String("account-id", "", "Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
		account_deleteAlternateContactCmd.Flags().String("alternate-contact-type", "", "Specifies which of the alternate contacts to delete.")
		account_deleteAlternateContactCmd.MarkFlagRequired("alternate-contact-type")
	})
	accountCmd.AddCommand(account_deleteAlternateContactCmd)
}
