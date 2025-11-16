package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_getPrimaryEmailCmd = &cobra.Command{
	Use:   "get-primary-email",
	Short: "Retrieves the primary email address for the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_getPrimaryEmailCmd).Standalone()

	account_getPrimaryEmailCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	account_getPrimaryEmailCmd.MarkFlagRequired("account-id")
	accountCmd.AddCommand(account_getPrimaryEmailCmd)
}
