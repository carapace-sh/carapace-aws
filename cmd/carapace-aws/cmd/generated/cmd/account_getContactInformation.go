package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_getContactInformationCmd = &cobra.Command{
	Use:   "get-contact-information",
	Short: "Retrieves the primary contact information of an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_getContactInformationCmd).Standalone()

	account_getContactInformationCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	accountCmd.AddCommand(account_getContactInformationCmd)
}
