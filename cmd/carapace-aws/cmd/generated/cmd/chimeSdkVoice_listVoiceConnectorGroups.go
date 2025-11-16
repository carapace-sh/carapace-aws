package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listVoiceConnectorGroupsCmd = &cobra.Command{
	Use:   "list-voice-connector-groups",
	Short: "Lists the Amazon Chime SDK Voice Connector groups in the administrator's AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listVoiceConnectorGroupsCmd).Standalone()

	chimeSdkVoice_listVoiceConnectorGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chimeSdkVoice_listVoiceConnectorGroupsCmd.Flags().String("next-token", "", "The token used to return the next page of results.")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listVoiceConnectorGroupsCmd)
}
