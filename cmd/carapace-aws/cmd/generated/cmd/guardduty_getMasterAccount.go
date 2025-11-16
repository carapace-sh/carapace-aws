package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getMasterAccountCmd = &cobra.Command{
	Use:   "get-master-account",
	Short: "Provides the details for the GuardDuty administrator account associated with the current GuardDuty member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getMasterAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getMasterAccountCmd).Standalone()

		guardduty_getMasterAccountCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty member account.")
		guardduty_getMasterAccountCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_getMasterAccountCmd)
}
