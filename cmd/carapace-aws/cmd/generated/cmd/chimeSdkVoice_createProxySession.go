package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createProxySessionCmd = &cobra.Command{
	Use:   "create-proxy-session",
	Short: "Creates a proxy session for the specified Amazon Chime SDK Voice Connector for the specified participant phone numbers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createProxySessionCmd).Standalone()

	chimeSdkVoice_createProxySessionCmd.Flags().String("capabilities", "", "The proxy session's capabilities.")
	chimeSdkVoice_createProxySessionCmd.Flags().String("expiry-minutes", "", "The number of minutes allowed for the proxy session.")
	chimeSdkVoice_createProxySessionCmd.Flags().String("geo-match-level", "", "The preference for matching the country or area code of the proxy phone number with that of the first participant.")
	chimeSdkVoice_createProxySessionCmd.Flags().String("geo-match-params", "", "The country and area code for the proxy phone number.")
	chimeSdkVoice_createProxySessionCmd.Flags().String("name", "", "The name of the proxy session.")
	chimeSdkVoice_createProxySessionCmd.Flags().String("number-selection-behavior", "", "The preference for proxy phone number reuse, or stickiness, between the same participants across sessions.")
	chimeSdkVoice_createProxySessionCmd.Flags().String("participant-phone-numbers", "", "The participant phone numbers.")
	chimeSdkVoice_createProxySessionCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_createProxySessionCmd.MarkFlagRequired("capabilities")
	chimeSdkVoice_createProxySessionCmd.MarkFlagRequired("participant-phone-numbers")
	chimeSdkVoice_createProxySessionCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createProxySessionCmd)
}
