package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_putAccountNameCmd = &cobra.Command{
	Use:   "put-account-name",
	Short: "Updates the account name of the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_putAccountNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(account_putAccountNameCmd).Standalone()

		account_putAccountNameCmd.Flags().String("account-id", "", "Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
		account_putAccountNameCmd.Flags().String("account-name", "", "The name of the account.")
		account_putAccountNameCmd.MarkFlagRequired("account-name")
	})
	accountCmd.AddCommand(account_putAccountNameCmd)
}
