package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeRegistrationAttachmentsCmd = &cobra.Command{
	Use:   "describe-registration-attachments",
	Short: "Retrieves the specified registration attachments or all registration attachments associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeRegistrationAttachmentsCmd).Standalone()

	pinpointSmsVoiceV2_describeRegistrationAttachmentsCmd.Flags().String("filters", "", "An array of RegistrationAttachmentFilter objects to filter the results.")
	pinpointSmsVoiceV2_describeRegistrationAttachmentsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeRegistrationAttachmentsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeRegistrationAttachmentsCmd.Flags().String("registration-attachment-ids", "", "The unique identifier of registration attachments to find.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeRegistrationAttachmentsCmd)
}
