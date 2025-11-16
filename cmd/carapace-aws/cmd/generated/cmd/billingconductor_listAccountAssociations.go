package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listAccountAssociationsCmd = &cobra.Command{
	Use:   "list-account-associations",
	Short: "This is a paginated call to list linked accounts that are linked to the payer account for the specified time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listAccountAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listAccountAssociationsCmd).Standalone()

		billingconductor_listAccountAssociationsCmd.Flags().String("billing-period", "", "The preferred billing period to get account associations.")
		billingconductor_listAccountAssociationsCmd.Flags().String("filters", "", "The filter on the account ID of the linked account, or any of the following:")
		billingconductor_listAccountAssociationsCmd.Flags().String("next-token", "", "The pagination token that's used on subsequent calls to retrieve accounts.")
	})
	billingconductorCmd.AddCommand(billingconductor_listAccountAssociationsCmd)
}
