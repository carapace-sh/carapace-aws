package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createRegistrationVersionCmd = &cobra.Command{
	Use:   "create-registration-version",
	Short: "Create a new version of the registration and increase the **VersionNumber**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createRegistrationVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_createRegistrationVersionCmd).Standalone()

		pinpointSmsVoiceV2_createRegistrationVersionCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
		pinpointSmsVoiceV2_createRegistrationVersionCmd.MarkFlagRequired("registration-id")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createRegistrationVersionCmd)
}
