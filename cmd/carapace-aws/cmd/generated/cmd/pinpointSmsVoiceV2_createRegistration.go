package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createRegistrationCmd = &cobra.Command{
	Use:   "create-registration",
	Short: "Creates a new registration based on the **RegistrationType** field.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_createRegistrationCmd).Standalone()

		pinpointSmsVoiceV2_createRegistrationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pinpointSmsVoiceV2_createRegistrationCmd.Flags().String("registration-type", "", "The type of registration form to create.")
		pinpointSmsVoiceV2_createRegistrationCmd.Flags().String("tags", "", "An array of tags (key and value pairs) to associate with the registration.")
		pinpointSmsVoiceV2_createRegistrationCmd.MarkFlagRequired("registration-type")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createRegistrationCmd)
}
