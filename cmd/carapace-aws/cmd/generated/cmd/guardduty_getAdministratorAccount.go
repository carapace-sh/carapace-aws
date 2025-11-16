package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getAdministratorAccountCmd = &cobra.Command{
	Use:   "get-administrator-account",
	Short: "Provides the details of the GuardDuty administrator account associated with the current GuardDuty member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getAdministratorAccountCmd).Standalone()

	guardduty_getAdministratorAccountCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty member account.")
	guardduty_getAdministratorAccountCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_getAdministratorAccountCmd)
}
