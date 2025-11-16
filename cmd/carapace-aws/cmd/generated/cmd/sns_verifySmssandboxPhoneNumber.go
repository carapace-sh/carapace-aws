package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_verifySmssandboxPhoneNumberCmd = &cobra.Command{
	Use:   "verify-smssandbox-phone-number",
	Short: "Verifies a destination phone number with a one-time password (OTP) for the calling Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_verifySmssandboxPhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_verifySmssandboxPhoneNumberCmd).Standalone()

		sns_verifySmssandboxPhoneNumberCmd.Flags().String("one-time-password", "", "The OTP sent to the destination number from the `CreateSMSSandBoxPhoneNumber` call.")
		sns_verifySmssandboxPhoneNumberCmd.Flags().String("phone-number", "", "The destination phone number to verify.")
		sns_verifySmssandboxPhoneNumberCmd.MarkFlagRequired("one-time-password")
		sns_verifySmssandboxPhoneNumberCmd.MarkFlagRequired("phone-number")
	})
	snsCmd.AddCommand(sns_verifySmssandboxPhoneNumberCmd)
}
