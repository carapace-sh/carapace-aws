package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_putAlternateContactCmd = &cobra.Command{
	Use:   "put-alternate-contact",
	Short: "Modifies the specified alternate contact attached to an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_putAlternateContactCmd).Standalone()

	account_putAlternateContactCmd.Flags().String("account-id", "", "Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	account_putAlternateContactCmd.Flags().String("alternate-contact-type", "", "Specifies which alternate contact you want to create or update.")
	account_putAlternateContactCmd.Flags().String("email-address", "", "Specifies an email address for the alternate contact.")
	account_putAlternateContactCmd.Flags().String("name", "", "Specifies a name for the alternate contact.")
	account_putAlternateContactCmd.Flags().String("phone-number", "", "Specifies a phone number for the alternate contact.")
	account_putAlternateContactCmd.Flags().String("title", "", "Specifies a title for the alternate contact.")
	account_putAlternateContactCmd.MarkFlagRequired("alternate-contact-type")
	account_putAlternateContactCmd.MarkFlagRequired("email-address")
	account_putAlternateContactCmd.MarkFlagRequired("name")
	account_putAlternateContactCmd.MarkFlagRequired("phone-number")
	account_putAlternateContactCmd.MarkFlagRequired("title")
	accountCmd.AddCommand(account_putAlternateContactCmd)
}
