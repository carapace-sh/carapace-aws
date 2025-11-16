package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteVerifiedDestinationNumberCmd = &cobra.Command{
	Use:   "delete-verified-destination-number",
	Short: "Delete a verified destination phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteVerifiedDestinationNumberCmd).Standalone()

	pinpointSmsVoiceV2_deleteVerifiedDestinationNumberCmd.Flags().String("verified-destination-number-id", "", "The unique identifier for the verified destination phone number.")
	pinpointSmsVoiceV2_deleteVerifiedDestinationNumberCmd.MarkFlagRequired("verified-destination-number-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteVerifiedDestinationNumberCmd)
}
