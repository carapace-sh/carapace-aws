package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createRegistrationAttachmentCmd = &cobra.Command{
	Use:   "create-registration-attachment",
	Short: "Create a new registration attachment to use for uploading a file or a URL to a file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createRegistrationAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_createRegistrationAttachmentCmd).Standalone()

		pinpointSmsVoiceV2_createRegistrationAttachmentCmd.Flags().String("attachment-body", "", "The registration file to upload.")
		pinpointSmsVoiceV2_createRegistrationAttachmentCmd.Flags().String("attachment-url", "", "Registration files have to be stored in an Amazon S3 bucket.")
		pinpointSmsVoiceV2_createRegistrationAttachmentCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pinpointSmsVoiceV2_createRegistrationAttachmentCmd.Flags().String("tags", "", "An array of tags (key and value pairs) to associate with the registration attachment.")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createRegistrationAttachmentCmd)
}
