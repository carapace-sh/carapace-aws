package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listBonusPaymentsCmd = &cobra.Command{
	Use:   "list-bonus-payments",
	Short: "The `ListBonusPayments` operation retrieves the amounts of bonuses you have paid to Workers for a given HIT or assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listBonusPaymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_listBonusPaymentsCmd).Standalone()

		mturk_listBonusPaymentsCmd.Flags().String("assignment-id", "", "The ID of the assignment associated with the bonus payments to retrieve.")
		mturk_listBonusPaymentsCmd.Flags().String("hitid", "", "The ID of the HIT associated with the bonus payments to retrieve.")
		mturk_listBonusPaymentsCmd.Flags().String("max-results", "", "")
		mturk_listBonusPaymentsCmd.Flags().String("next-token", "", "Pagination token")
	})
	mturkCmd.AddCommand(mturk_listBonusPaymentsCmd)
}
