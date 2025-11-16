package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_listRegistrationAssociationsCmd = &cobra.Command{
	Use:   "list-registration-associations",
	Short: "Retrieve all of the origination identities that are associated with a registration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_listRegistrationAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_listRegistrationAssociationsCmd).Standalone()

		pinpointSmsVoiceV2_listRegistrationAssociationsCmd.Flags().String("filters", "", "An array of RegistrationAssociationFilter to apply to the results that are returned.")
		pinpointSmsVoiceV2_listRegistrationAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per each request.")
		pinpointSmsVoiceV2_listRegistrationAssociationsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		pinpointSmsVoiceV2_listRegistrationAssociationsCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
		pinpointSmsVoiceV2_listRegistrationAssociationsCmd.MarkFlagRequired("registration-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_listRegistrationAssociationsCmd)
}
