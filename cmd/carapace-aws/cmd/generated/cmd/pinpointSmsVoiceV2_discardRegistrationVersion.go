package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_discardRegistrationVersionCmd = &cobra.Command{
	Use:   "discard-registration-version",
	Short: "Discard the current version of the registration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_discardRegistrationVersionCmd).Standalone()

	pinpointSmsVoiceV2_discardRegistrationVersionCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
	pinpointSmsVoiceV2_discardRegistrationVersionCmd.MarkFlagRequired("registration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_discardRegistrationVersionCmd)
}
