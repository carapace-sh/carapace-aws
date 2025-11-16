package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_sendContactMethodVerificationCmd = &cobra.Command{
	Use:   "send-contact-method-verification",
	Short: "Sends a verification request to an email contact method to ensure it's owned by the requester.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_sendContactMethodVerificationCmd).Standalone()

	lightsail_sendContactMethodVerificationCmd.Flags().String("protocol", "", "The protocol to verify, such as `Email` or `SMS` (text messaging).")
	lightsail_sendContactMethodVerificationCmd.MarkFlagRequired("protocol")
	lightsailCmd.AddCommand(lightsail_sendContactMethodVerificationCmd)
}
