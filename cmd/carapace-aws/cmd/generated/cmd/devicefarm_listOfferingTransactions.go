package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listOfferingTransactionsCmd = &cobra.Command{
	Use:   "list-offering-transactions",
	Short: "Returns a list of all historical purchases, renewals, and system renewal transactions for an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listOfferingTransactionsCmd).Standalone()

	devicefarm_listOfferingTransactionsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarmCmd.AddCommand(devicefarm_listOfferingTransactionsCmd)
}
