package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeRegistrationVersionsCmd = &cobra.Command{
	Use:   "describe-registration-versions",
	Short: "Retrieves the specified registration version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeRegistrationVersionsCmd).Standalone()

	pinpointSmsVoiceV2_describeRegistrationVersionsCmd.Flags().String("filters", "", "An array of RegistrationVersionFilter objects to filter the results.")
	pinpointSmsVoiceV2_describeRegistrationVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeRegistrationVersionsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeRegistrationVersionsCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
	pinpointSmsVoiceV2_describeRegistrationVersionsCmd.Flags().String("version-numbers", "", "An array of registration version numbers.")
	pinpointSmsVoiceV2_describeRegistrationVersionsCmd.MarkFlagRequired("registration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeRegistrationVersionsCmd)
}
