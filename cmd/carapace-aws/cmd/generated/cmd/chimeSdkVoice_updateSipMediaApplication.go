package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateSipMediaApplicationCmd = &cobra.Command{
	Use:   "update-sip-media-application",
	Short: "Updates the details of the specified SIP media application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateSipMediaApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_updateSipMediaApplicationCmd).Standalone()

		chimeSdkVoice_updateSipMediaApplicationCmd.Flags().String("endpoints", "", "The new set of endpoints for the specified SIP media application.")
		chimeSdkVoice_updateSipMediaApplicationCmd.Flags().String("name", "", "The new name for the specified SIP media application.")
		chimeSdkVoice_updateSipMediaApplicationCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID.")
		chimeSdkVoice_updateSipMediaApplicationCmd.MarkFlagRequired("sip-media-application-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateSipMediaApplicationCmd)
}
