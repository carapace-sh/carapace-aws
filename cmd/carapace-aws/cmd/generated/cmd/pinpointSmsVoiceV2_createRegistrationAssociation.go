package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createRegistrationAssociationCmd = &cobra.Command{
	Use:   "create-registration-association",
	Short: "Associate the registration with an origination identity such as a phone number or sender ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createRegistrationAssociationCmd).Standalone()

	pinpointSmsVoiceV2_createRegistrationAssociationCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
	pinpointSmsVoiceV2_createRegistrationAssociationCmd.Flags().String("resource-id", "", "The unique identifier for the origination identity.")
	pinpointSmsVoiceV2_createRegistrationAssociationCmd.MarkFlagRequired("registration-id")
	pinpointSmsVoiceV2_createRegistrationAssociationCmd.MarkFlagRequired("resource-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createRegistrationAssociationCmd)
}
