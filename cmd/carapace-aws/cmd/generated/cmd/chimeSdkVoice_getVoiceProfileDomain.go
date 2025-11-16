package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceProfileDomainCmd = &cobra.Command{
	Use:   "get-voice-profile-domain",
	Short: "Retrieves the details of the specified voice profile domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceProfileDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getVoiceProfileDomainCmd).Standalone()

		chimeSdkVoice_getVoiceProfileDomainCmd.Flags().String("voice-profile-domain-id", "", "The voice profile domain ID.")
		chimeSdkVoice_getVoiceProfileDomainCmd.MarkFlagRequired("voice-profile-domain-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceProfileDomainCmd)
}
