package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteRegistrationFieldValueCmd = &cobra.Command{
	Use:   "delete-registration-field-value",
	Short: "Delete the value in a registration form field.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteRegistrationFieldValueCmd).Standalone()

	pinpointSmsVoiceV2_deleteRegistrationFieldValueCmd.Flags().String("field-path", "", "The path to the registration form field.")
	pinpointSmsVoiceV2_deleteRegistrationFieldValueCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
	pinpointSmsVoiceV2_deleteRegistrationFieldValueCmd.MarkFlagRequired("field-path")
	pinpointSmsVoiceV2_deleteRegistrationFieldValueCmd.MarkFlagRequired("registration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteRegistrationFieldValueCmd)
}
