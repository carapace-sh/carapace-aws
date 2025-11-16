package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteRegistrationCmd = &cobra.Command{
	Use:   "delete-registration",
	Short: "Permanently delete an existing registration from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_deleteRegistrationCmd).Standalone()

		pinpointSmsVoiceV2_deleteRegistrationCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
		pinpointSmsVoiceV2_deleteRegistrationCmd.MarkFlagRequired("registration-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteRegistrationCmd)
}
