package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_getAlternateContactCmd = &cobra.Command{
	Use:   "get-alternate-contact",
	Short: "Retrieves the specified alternate contact attached to an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_getAlternateContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(account_getAlternateContactCmd).Standalone()

		account_getAlternateContactCmd.Flags().String("account-id", "", "Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
		account_getAlternateContactCmd.Flags().String("alternate-contact-type", "", "Specifies which alternate contact you want to retrieve.")
		account_getAlternateContactCmd.MarkFlagRequired("alternate-contact-type")
	})
	accountCmd.AddCommand(account_getAlternateContactCmd)
}
