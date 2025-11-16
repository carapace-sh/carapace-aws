package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteSipMediaApplicationCmd = &cobra.Command{
	Use:   "delete-sip-media-application",
	Short: "Deletes a SIP media application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteSipMediaApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deleteSipMediaApplicationCmd).Standalone()

		chimeSdkVoice_deleteSipMediaApplicationCmd.Flags().String("sip-media-application-id", "", "The SIP media application ID.")
		chimeSdkVoice_deleteSipMediaApplicationCmd.MarkFlagRequired("sip-media-application-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteSipMediaApplicationCmd)
}
