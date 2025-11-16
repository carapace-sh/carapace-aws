package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listVoiceProfilesCmd = &cobra.Command{
	Use:   "list-voice-profiles",
	Short: "Lists the voice profiles in a voice profile domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listVoiceProfilesCmd).Standalone()

	chimeSdkVoice_listVoiceProfilesCmd.Flags().String("max-results", "", "The maximum number of results in the request.")
	chimeSdkVoice_listVoiceProfilesCmd.Flags().String("next-token", "", "The token used to retrieve the next page of results.")
	chimeSdkVoice_listVoiceProfilesCmd.Flags().String("voice-profile-domain-id", "", "The ID of the voice profile domain.")
	chimeSdkVoice_listVoiceProfilesCmd.MarkFlagRequired("voice-profile-domain-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listVoiceProfilesCmd)
}
