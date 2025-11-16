package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listVoiceConnectorsCmd = &cobra.Command{
	Use:   "list-voice-connectors",
	Short: "Lists the Amazon Chime SDK Voice Connectors in the administrators AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listVoiceConnectorsCmd).Standalone()

	chimeSdkVoice_listVoiceConnectorsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chimeSdkVoice_listVoiceConnectorsCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listVoiceConnectorsCmd)
}
