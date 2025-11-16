package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_describeRegistrationsCmd = &cobra.Command{
	Use:   "describe-registrations",
	Short: "Retrieves the specified registrations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_describeRegistrationsCmd).Standalone()

	pinpointSmsVoiceV2_describeRegistrationsCmd.Flags().String("filters", "", "An array of RegistrationFilter objects to filter the results.")
	pinpointSmsVoiceV2_describeRegistrationsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
	pinpointSmsVoiceV2_describeRegistrationsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	pinpointSmsVoiceV2_describeRegistrationsCmd.Flags().String("registration-ids", "", "An array of unique identifiers for each registration.")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_describeRegistrationsCmd)
}
