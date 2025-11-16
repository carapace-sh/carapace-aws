package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createVerifiedDestinationNumberCmd = &cobra.Command{
	Use:   "create-verified-destination-number",
	Short: "You can only send messages to verified destination numbers when your account is in the sandbox.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createVerifiedDestinationNumberCmd).Standalone()

	pinpointSmsVoiceV2_createVerifiedDestinationNumberCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pinpointSmsVoiceV2_createVerifiedDestinationNumberCmd.Flags().String("destination-phone-number", "", "The verified destination phone number, in E.164 format.")
	pinpointSmsVoiceV2_createVerifiedDestinationNumberCmd.Flags().String("tags", "", "An array of tags (key and value pairs) to associate with the destination number.")
	pinpointSmsVoiceV2_createVerifiedDestinationNumberCmd.MarkFlagRequired("destination-phone-number")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createVerifiedDestinationNumberCmd)
}
