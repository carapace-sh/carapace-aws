package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoSignaling_sendAlexaOfferToMasterCmd = &cobra.Command{
	Use:   "send-alexa-offer-to-master",
	Short: "This API allows you to connect WebRTC-enabled devices with Alexa display devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoSignaling_sendAlexaOfferToMasterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoSignaling_sendAlexaOfferToMasterCmd).Standalone()

		kinesisVideoSignaling_sendAlexaOfferToMasterCmd.Flags().String("channel-arn", "", "The ARN of the signaling channel by which Alexa and the master peer communicate.")
		kinesisVideoSignaling_sendAlexaOfferToMasterCmd.Flags().String("message-payload", "", "The base64-encoded SDP offer content.")
		kinesisVideoSignaling_sendAlexaOfferToMasterCmd.Flags().String("sender-client-id", "", "The unique identifier for the sender client.")
		kinesisVideoSignaling_sendAlexaOfferToMasterCmd.MarkFlagRequired("channel-arn")
		kinesisVideoSignaling_sendAlexaOfferToMasterCmd.MarkFlagRequired("message-payload")
		kinesisVideoSignaling_sendAlexaOfferToMasterCmd.MarkFlagRequired("sender-client-id")
	})
	kinesisVideoSignalingCmd.AddCommand(kinesisVideoSignaling_sendAlexaOfferToMasterCmd)
}
