package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_disassociateFromMasterAccountCmd = &cobra.Command{
	Use:   "disassociate-from-master-account",
	Short: "Disassociates the current GuardDuty member account from its administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_disassociateFromMasterAccountCmd).Standalone()

	guardduty_disassociateFromMasterAccountCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty member account.")
	guardduty_disassociateFromMasterAccountCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_disassociateFromMasterAccountCmd)
}
