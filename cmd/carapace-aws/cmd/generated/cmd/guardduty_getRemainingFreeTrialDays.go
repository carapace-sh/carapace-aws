package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getRemainingFreeTrialDaysCmd = &cobra.Command{
	Use:   "get-remaining-free-trial-days",
	Short: "Provides the number of days left for each data source used in the free trial period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getRemainingFreeTrialDaysCmd).Standalone()

	guardduty_getRemainingFreeTrialDaysCmd.Flags().String("account-ids", "", "A list of account identifiers of the GuardDuty member account.")
	guardduty_getRemainingFreeTrialDaysCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty member account.")
	guardduty_getRemainingFreeTrialDaysCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_getRemainingFreeTrialDaysCmd)
}
