package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeRegistrationTypeDefinitionsCmd = &cobra.Command{
	Use:   "describe-registration-type-definitions",
	Short: "Retrieves the specified registration type definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeRegistrationTypeDefinitionsCmd).Standalone()

	pinpointSmsVoiceV2_describeRegistrationTypeDefinitionsCmd.Flags().String("filters", "", "An array of RegistrationFilter objects to filter the results.")
	pinpointSmsVoiceV2_describeRegistrationTypeDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeRegistrationTypeDefinitionsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeRegistrationTypeDefinitionsCmd.Flags().String("registration-types", "", "The type of registration form.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeRegistrationTypeDefinitionsCmd)
}
