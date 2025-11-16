package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_disassociateFromAdministratorAccountCmd = &cobra.Command{
	Use:   "disassociate-from-administrator-account",
	Short: "Disassociates the current GuardDuty member account from its administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_disassociateFromAdministratorAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_disassociateFromAdministratorAccountCmd).Standalone()

		guardduty_disassociateFromAdministratorAccountCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty member account.")
		guardduty_disassociateFromAdministratorAccountCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_disassociateFromAdministratorAccountCmd)
}
