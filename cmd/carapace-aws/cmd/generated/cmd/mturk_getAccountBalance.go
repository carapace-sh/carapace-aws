package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_getAccountBalanceCmd = &cobra.Command{
	Use:   "get-account-balance",
	Short: "The `GetAccountBalance` operation retrieves the Prepaid HITs balance in your Amazon Mechanical Turk account if you are a Prepaid Requester.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_getAccountBalanceCmd).Standalone()

	mturkCmd.AddCommand(mturk_getAccountBalanceCmd)
}
