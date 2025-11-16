package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceProfileDomainCmd = &cobra.Command{
	Use:   "delete-voice-profile-domain",
	Short: "Deletes all voice profiles in the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceProfileDomainCmd).Standalone()

	chimeSdkVoice_deleteVoiceProfileDomainCmd.Flags().String("voice-profile-domain-id", "", "The voice profile domain ID.")
	chimeSdkVoice_deleteVoiceProfileDomainCmd.MarkFlagRequired("voice-profile-domain-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceProfileDomainCmd)
}
