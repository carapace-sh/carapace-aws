package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_putContactInformationCmd = &cobra.Command{
	Use:   "put-contact-information",
	Short: "Updates the primary contact information of an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_putContactInformationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(account_putContactInformationCmd).Standalone()

		account_putContactInformationCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
		account_putContactInformationCmd.Flags().String("contact-information", "", "Contains the details of the primary contact information associated with an Amazon Web Services account.")
		account_putContactInformationCmd.MarkFlagRequired("contact-information")
	})
	accountCmd.AddCommand(account_putContactInformationCmd)
}
