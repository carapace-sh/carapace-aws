package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_associateAccountsCmd = &cobra.Command{
	Use:   "associate-accounts",
	Short: "Connects an array of account IDs in a consolidated billing family to a predefined billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_associateAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_associateAccountsCmd).Standalone()

		billingconductor_associateAccountsCmd.Flags().String("account-ids", "", "The associating array of account IDs.")
		billingconductor_associateAccountsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the billing group that associates the array of account IDs.")
		billingconductor_associateAccountsCmd.MarkFlagRequired("account-ids")
		billingconductor_associateAccountsCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_associateAccountsCmd)
}
