package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_getAccountInformationCmd = &cobra.Command{
	Use:   "get-account-information",
	Short: "Retrieves information about the specified account including its account name, account ID, and account creation date and time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_getAccountInformationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(account_getAccountInformationCmd).Standalone()

		account_getAccountInformationCmd.Flags().String("account-id", "", "Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	})
	accountCmd.AddCommand(account_getAccountInformationCmd)
}
