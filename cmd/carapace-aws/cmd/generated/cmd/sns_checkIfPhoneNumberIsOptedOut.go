package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_checkIfPhoneNumberIsOptedOutCmd = &cobra.Command{
	Use:   "check-if-phone-number-is-opted-out",
	Short: "Accepts a phone number and indicates whether the phone holder has opted out of receiving SMS messages from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_checkIfPhoneNumberIsOptedOutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_checkIfPhoneNumberIsOptedOutCmd).Standalone()

		sns_checkIfPhoneNumberIsOptedOutCmd.Flags().String("phone-number", "", "The phone number for which you want to check the opt out status.")
		sns_checkIfPhoneNumberIsOptedOutCmd.MarkFlagRequired("phone-number")
	})
	snsCmd.AddCommand(sns_checkIfPhoneNumberIsOptedOutCmd)
}
