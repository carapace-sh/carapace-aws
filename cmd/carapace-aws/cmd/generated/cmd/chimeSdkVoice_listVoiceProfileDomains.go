package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listVoiceProfileDomainsCmd = &cobra.Command{
	Use:   "list-voice-profile-domains",
	Short: "Lists the specified voice profile domains in the administrator's AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listVoiceProfileDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_listVoiceProfileDomainsCmd).Standalone()

		chimeSdkVoice_listVoiceProfileDomainsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chimeSdkVoice_listVoiceProfileDomainsCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listVoiceProfileDomainsCmd)
}
