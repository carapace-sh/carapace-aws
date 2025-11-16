package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd = &cobra.Command{
	Use:   "describe-registration-field-definitions",
	Short: "Retrieves the specified registration type field definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd).Standalone()

	pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd.Flags().String("field-paths", "", "An array of paths to the registration form field.")
	pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd.Flags().String("registration-type", "", "The type of registration form.")
	pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd.Flags().String("section-path", "", "The path to the section of the registration.")
	pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd.MarkFlagRequired("registration-type")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeRegistrationFieldDefinitionsCmd)
}
