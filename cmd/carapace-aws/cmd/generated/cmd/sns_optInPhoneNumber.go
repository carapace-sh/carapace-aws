package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_optInPhoneNumberCmd = &cobra.Command{
	Use:   "opt-in-phone-number",
	Short: "Use this request to opt in a phone number that is opted out, which enables you to resume sending SMS messages to the number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_optInPhoneNumberCmd).Standalone()

	sns_optInPhoneNumberCmd.Flags().String("phone-number", "", "The phone number to opt in.")
	sns_optInPhoneNumberCmd.MarkFlagRequired("phone-number")
	snsCmd.AddCommand(sns_optInPhoneNumberCmd)
}
