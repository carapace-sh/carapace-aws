package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getPhoneNumberSettingsCmd = &cobra.Command{
	Use:   "get-phone-number-settings",
	Short: "Retrieves the phone number settings for the administrator's AWS account, such as the default outbound calling name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getPhoneNumberSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getPhoneNumberSettingsCmd).Standalone()

	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getPhoneNumberSettingsCmd)
}
