package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_sendBonusCmd = &cobra.Command{
	Use:   "send-bonus",
	Short: "The `SendBonus` operation issues a payment of money from your account to a Worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_sendBonusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_sendBonusCmd).Standalone()

		mturk_sendBonusCmd.Flags().String("assignment-id", "", "The ID of the assignment for which this bonus is paid.")
		mturk_sendBonusCmd.Flags().String("bonus-amount", "", "The Bonus amount is a US Dollar amount specified using a string (for example, \"5\" represents $5.00 USD and \"101.42\" represents $101.42 USD).")
		mturk_sendBonusCmd.Flags().String("reason", "", "A message that explains the reason for the bonus payment.")
		mturk_sendBonusCmd.Flags().String("unique-request-token", "", "A unique identifier for this request, which allows you to retry the call on error without granting multiple bonuses.")
		mturk_sendBonusCmd.Flags().String("worker-id", "", "The ID of the Worker being paid the bonus.")
		mturk_sendBonusCmd.MarkFlagRequired("assignment-id")
		mturk_sendBonusCmd.MarkFlagRequired("bonus-amount")
		mturk_sendBonusCmd.MarkFlagRequired("reason")
		mturk_sendBonusCmd.MarkFlagRequired("worker-id")
	})
	mturkCmd.AddCommand(mturk_sendBonusCmd)
}
