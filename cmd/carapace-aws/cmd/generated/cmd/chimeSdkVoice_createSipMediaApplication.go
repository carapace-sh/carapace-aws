package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createSipMediaApplicationCmd = &cobra.Command{
	Use:   "create-sip-media-application",
	Short: "Creates a SIP media application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createSipMediaApplicationCmd).Standalone()

	chimeSdkVoice_createSipMediaApplicationCmd.Flags().String("aws-region", "", "The AWS Region assigned to the SIP media application.")
	chimeSdkVoice_createSipMediaApplicationCmd.Flags().String("endpoints", "", "List of endpoints (Lambda ARNs) specified for the SIP media application.")
	chimeSdkVoice_createSipMediaApplicationCmd.Flags().String("name", "", "The SIP media application's name.")
	chimeSdkVoice_createSipMediaApplicationCmd.Flags().String("tags", "", "The tags assigned to the SIP media application.")
	chimeSdkVoice_createSipMediaApplicationCmd.MarkFlagRequired("aws-region")
	chimeSdkVoice_createSipMediaApplicationCmd.MarkFlagRequired("endpoints")
	chimeSdkVoice_createSipMediaApplicationCmd.MarkFlagRequired("name")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createSipMediaApplicationCmd)
}
