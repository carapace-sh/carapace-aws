package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createVoiceConnectorGroupCmd = &cobra.Command{
	Use:   "create-voice-connector-group",
	Short: "Creates an Amazon Chime SDK Voice Connector group under the administrator's AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createVoiceConnectorGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_createVoiceConnectorGroupCmd).Standalone()

		chimeSdkVoice_createVoiceConnectorGroupCmd.Flags().String("name", "", "The name of the Voice Connector group.")
		chimeSdkVoice_createVoiceConnectorGroupCmd.Flags().String("voice-connector-items", "", "Lists the Voice Connectors that inbound calls are routed to.")
		chimeSdkVoice_createVoiceConnectorGroupCmd.MarkFlagRequired("name")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createVoiceConnectorGroupCmd)
}
