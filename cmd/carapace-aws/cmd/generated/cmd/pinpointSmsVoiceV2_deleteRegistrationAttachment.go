package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteRegistrationAttachmentCmd = &cobra.Command{
	Use:   "delete-registration-attachment",
	Short: "Permanently delete the specified registration attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteRegistrationAttachmentCmd).Standalone()

	pinpointSmsVoiceV2_deleteRegistrationAttachmentCmd.Flags().String("registration-attachment-id", "", "The unique identifier for the registration attachment.")
	pinpointSmsVoiceV2_deleteRegistrationAttachmentCmd.MarkFlagRequired("registration-attachment-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteRegistrationAttachmentCmd)
}
