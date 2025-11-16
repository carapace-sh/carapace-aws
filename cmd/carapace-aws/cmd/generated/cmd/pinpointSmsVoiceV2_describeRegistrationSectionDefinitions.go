package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd = &cobra.Command{
	Use:   "describe-registration-section-definitions",
	Short: "Retrieves the specified registration section definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd).Standalone()

		pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd.Flags().String("registration-type", "", "The type of registration form.")
		pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd.Flags().String("section-paths", "", "An array of paths for the registration form section.")
		pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd.MarkFlagRequired("registration-type")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeRegistrationSectionDefinitionsCmd)
}
