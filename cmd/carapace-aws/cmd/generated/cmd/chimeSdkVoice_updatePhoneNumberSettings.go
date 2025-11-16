package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updatePhoneNumberSettingsCmd = &cobra.Command{
	Use:   "update-phone-number-settings",
	Short: "Updates the phone number settings for the administrator's AWS account, such as the default outbound calling name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updatePhoneNumberSettingsCmd).Standalone()

	chimeSdkVoice_updatePhoneNumberSettingsCmd.Flags().String("calling-name", "", "The default outbound calling name for the account.")
	chimeSdkVoice_updatePhoneNumberSettingsCmd.MarkFlagRequired("calling-name")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updatePhoneNumberSettingsCmd)
}
