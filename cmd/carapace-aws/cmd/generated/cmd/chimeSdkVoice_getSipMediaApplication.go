package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getSipMediaApplicationCmd = &cobra.Command{
	Use:   "get-sip-media-application",
	Short: "Retrieves the information for a SIP media application, including name, AWS Region, and endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getSipMediaApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_getSipMediaApplicationCmd).Standalone()

		chimeSdkVoice_getSipMediaApplicationCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID .")
		chimeSdkVoice_getSipMediaApplicationCmd.MarkFlagRequired("sip-media-application-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getSipMediaApplicationCmd)
}
