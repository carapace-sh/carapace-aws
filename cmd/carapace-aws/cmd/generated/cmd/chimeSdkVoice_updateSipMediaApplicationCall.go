package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateSipMediaApplicationCallCmd = &cobra.Command{
	Use:   "update-sip-media-application-call",
	Short: "Invokes the AWS Lambda function associated with the SIP media application and transaction ID in an update request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateSipMediaApplicationCallCmd).Standalone()

	chimeSdkVoice_updateSipMediaApplicationCallCmd.Flags().String("arguments", "", "Arguments made available to the Lambda function as part of the `CALL_UPDATE_REQUESTED` event.")
	chimeSdkVoice_updateSipMediaApplicationCallCmd.Flags().String("sip-media-application-id", "", "The ID of the SIP media application handling the call.")
	chimeSdkVoice_updateSipMediaApplicationCallCmd.Flags().String("transaction-id", "", "The ID of the call transaction.")
	chimeSdkVoice_updateSipMediaApplicationCallCmd.MarkFlagRequired("arguments")
	chimeSdkVoice_updateSipMediaApplicationCallCmd.MarkFlagRequired("sip-media-application-id")
	chimeSdkVoice_updateSipMediaApplicationCallCmd.MarkFlagRequired("transaction-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateSipMediaApplicationCallCmd)
}
