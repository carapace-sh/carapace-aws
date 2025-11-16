package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createVoiceProfileDomainCmd = &cobra.Command{
	Use:   "create-voice-profile-domain",
	Short: "Creates a voice profile domain, a collection of voice profiles, their voice prints, and encrypted enrollment audio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createVoiceProfileDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_createVoiceProfileDomainCmd).Standalone()

		chimeSdkVoice_createVoiceProfileDomainCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
		chimeSdkVoice_createVoiceProfileDomainCmd.Flags().String("description", "", "A description of the voice profile domain.")
		chimeSdkVoice_createVoiceProfileDomainCmd.Flags().String("name", "", "The name of the voice profile domain.")
		chimeSdkVoice_createVoiceProfileDomainCmd.Flags().String("server-side-encryption-configuration", "", "The server-side encryption configuration for the request.")
		chimeSdkVoice_createVoiceProfileDomainCmd.Flags().String("tags", "", "The tags assigned to the domain.")
		chimeSdkVoice_createVoiceProfileDomainCmd.MarkFlagRequired("name")
		chimeSdkVoice_createVoiceProfileDomainCmd.MarkFlagRequired("server-side-encryption-configuration")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createVoiceProfileDomainCmd)
}
