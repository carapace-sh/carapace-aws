package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorProxyCmd = &cobra.Command{
	Use:   "put-voice-connector-proxy",
	Short: "Puts the specified proxy configuration to the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorProxyCmd).Standalone()

	chimeSdkVoice_putVoiceConnectorProxyCmd.Flags().String("default-session-expiry-minutes", "", "The default number of minutes allowed for proxy session.")
	chimeSdkVoice_putVoiceConnectorProxyCmd.Flags().Bool("disabled", false, "When true, stops proxy sessions from being created on the specified Amazon Chime SDK Voice Connector.")
	chimeSdkVoice_putVoiceConnectorProxyCmd.Flags().String("fall-back-phone-number", "", "The phone number to route calls to after a proxy session expires.")
	chimeSdkVoice_putVoiceConnectorProxyCmd.Flags().Bool("no-disabled", false, "When true, stops proxy sessions from being created on the specified Amazon Chime SDK Voice Connector.")
	chimeSdkVoice_putVoiceConnectorProxyCmd.Flags().String("phone-number-pool-countries", "", "The countries for proxy phone numbers to be selected from.")
	chimeSdkVoice_putVoiceConnectorProxyCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_putVoiceConnectorProxyCmd.MarkFlagRequired("default-session-expiry-minutes")
	chimeSdkVoice_putVoiceConnectorProxyCmd.Flag("no-disabled").Hidden = true
	chimeSdkVoice_putVoiceConnectorProxyCmd.MarkFlagRequired("phone-number-pool-countries")
	chimeSdkVoice_putVoiceConnectorProxyCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorProxyCmd)
}
