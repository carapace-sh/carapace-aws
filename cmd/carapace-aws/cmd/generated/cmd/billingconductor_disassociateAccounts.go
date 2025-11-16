package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_disassociateAccountsCmd = &cobra.Command{
	Use:   "disassociate-accounts",
	Short: "Removes the specified list of account IDs from the given billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_disassociateAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_disassociateAccountsCmd).Standalone()

		billingconductor_disassociateAccountsCmd.Flags().String("account-ids", "", "The array of account IDs to disassociate.")
		billingconductor_disassociateAccountsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the billing group that the array of account IDs will disassociate from.")
		billingconductor_disassociateAccountsCmd.MarkFlagRequired("account-ids")
		billingconductor_disassociateAccountsCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_disassociateAccountsCmd)
}
