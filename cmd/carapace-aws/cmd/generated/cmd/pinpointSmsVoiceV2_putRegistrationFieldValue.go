package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_putRegistrationFieldValueCmd = &cobra.Command{
	Use:   "put-registration-field-value",
	Short: "Creates or updates a field value for a registration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_putRegistrationFieldValueCmd).Standalone()

	pinpointSmsVoiceV2_putRegistrationFieldValueCmd.Flags().String("field-path", "", "The path to the registration form field.")
	pinpointSmsVoiceV2_putRegistrationFieldValueCmd.Flags().String("registration-attachment-id", "", "The unique identifier for the registration attachment.")
	pinpointSmsVoiceV2_putRegistrationFieldValueCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
	pinpointSmsVoiceV2_putRegistrationFieldValueCmd.Flags().String("select-choices", "", "An array of values for the form field.")
	pinpointSmsVoiceV2_putRegistrationFieldValueCmd.Flags().String("text-value", "", "The text data for a free form field.")
	pinpointSmsVoiceV2_putRegistrationFieldValueCmd.MarkFlagRequired("field-path")
	pinpointSmsVoiceV2_putRegistrationFieldValueCmd.MarkFlagRequired("registration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_putRegistrationFieldValueCmd)
}
