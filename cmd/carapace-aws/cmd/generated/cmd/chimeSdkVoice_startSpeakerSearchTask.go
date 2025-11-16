package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_startSpeakerSearchTaskCmd = &cobra.Command{
	Use:   "start-speaker-search-task",
	Short: "Starts a speaker search task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_startSpeakerSearchTaskCmd).Standalone()

	chimeSdkVoice_startSpeakerSearchTaskCmd.Flags().String("call-leg", "", "Specifies which call leg to stream for speaker search.")
	chimeSdkVoice_startSpeakerSearchTaskCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
	chimeSdkVoice_startSpeakerSearchTaskCmd.Flags().String("transaction-id", "", "The transaction ID of the call being analyzed.")
	chimeSdkVoice_startSpeakerSearchTaskCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_startSpeakerSearchTaskCmd.Flags().String("voice-profile-domain-id", "", "The ID of the voice profile domain that will store the voice profile.")
	chimeSdkVoice_startSpeakerSearchTaskCmd.MarkFlagRequired("transaction-id")
	chimeSdkVoice_startSpeakerSearchTaskCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoice_startSpeakerSearchTaskCmd.MarkFlagRequired("voice-profile-domain-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_startSpeakerSearchTaskCmd)
}
