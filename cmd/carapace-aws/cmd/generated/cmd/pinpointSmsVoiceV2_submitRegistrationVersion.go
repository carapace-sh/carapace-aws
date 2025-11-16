package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_submitRegistrationVersionCmd = &cobra.Command{
	Use:   "submit-registration-version",
	Short: "Submit the specified registration for review and approval.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_submitRegistrationVersionCmd).Standalone()

	pinpointSmsVoiceV2_submitRegistrationVersionCmd.Flags().String("registration-id", "", "The unique identifier for the registration.")
	pinpointSmsVoiceV2_submitRegistrationVersionCmd.MarkFlagRequired("registration-id")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_submitRegistrationVersionCmd)
}
