package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_createSmssandboxPhoneNumberCmd = &cobra.Command{
	Use:   "create-smssandbox-phone-number",
	Short: "Adds a destination phone number to an Amazon Web Services account in the SMS sandbox and sends a one-time password (OTP) to that phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_createSmssandboxPhoneNumberCmd).Standalone()

	sns_createSmssandboxPhoneNumberCmd.Flags().String("language-code", "", "The language to use for sending the OTP.")
	sns_createSmssandboxPhoneNumberCmd.Flags().String("phone-number", "", "The destination phone number to verify.")
	sns_createSmssandboxPhoneNumberCmd.MarkFlagRequired("phone-number")
	snsCmd.AddCommand(sns_createSmssandboxPhoneNumberCmd)
}
