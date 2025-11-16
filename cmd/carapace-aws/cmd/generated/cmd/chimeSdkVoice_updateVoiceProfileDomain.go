package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateVoiceProfileDomainCmd = &cobra.Command{
	Use:   "update-voice-profile-domain",
	Short: "Updates the settings for the specified voice profile domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateVoiceProfileDomainCmd).Standalone()

	chimeSdkVoice_updateVoiceProfileDomainCmd.Flags().String("description", "", "The description of the voice profile domain.")
	chimeSdkVoice_updateVoiceProfileDomainCmd.Flags().String("name", "", "The name of the voice profile domain.")
	chimeSdkVoice_updateVoiceProfileDomainCmd.Flags().String("voice-profile-domain-id", "", "The domain ID.")
	chimeSdkVoice_updateVoiceProfileDomainCmd.MarkFlagRequired("voice-profile-domain-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateVoiceProfileDomainCmd)
}
