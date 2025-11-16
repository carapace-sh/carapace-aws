package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd = &cobra.Command{
	Use:   "describe-registration-field-values",
	Short: "Retrieves the specified registration field values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd).Standalone()

	pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd.Flags().String("field-paths", "", "An array of paths to the registration form field.")
	pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
	pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd.Flags().String("section-path", "", "The path to the section of the registration.")
	pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd.Flags().String("version-number", "", "The version number of the registration.")
	pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd.MarkFlagRequired("registration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeRegistrationFieldValuesCmd)
}
