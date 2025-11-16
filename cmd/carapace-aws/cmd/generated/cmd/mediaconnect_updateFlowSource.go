package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateFlowSourceCmd = &cobra.Command{
	Use:   "update-flow-source",
	Short: "Updates the source of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateFlowSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_updateFlowSourceCmd).Standalone()

		mediaconnect_updateFlowSourceCmd.Flags().String("decryption", "", "The type of encryption that is used on the content ingested from the source.")
		mediaconnect_updateFlowSourceCmd.Flags().String("description", "", "A description of the source.")
		mediaconnect_updateFlowSourceCmd.Flags().String("entitlement-arn", "", "The Amazon Resource Name (ARN) of the entitlement that allows you to subscribe to the flow.")
		mediaconnect_updateFlowSourceCmd.Flags().String("flow-arn", "", "The ARN of the flow that you want to update.")
		mediaconnect_updateFlowSourceCmd.Flags().String("gateway-bridge-source", "", "The source configuration for cloud flows receiving a stream from a bridge.")
		mediaconnect_updateFlowSourceCmd.Flags().String("ingest-port", "", "The port that the flow listens on for incoming content.")
		mediaconnect_updateFlowSourceCmd.Flags().String("max-bitrate", "", "The maximum bitrate for RIST, RTP, and RTP-FEC streams.")
		mediaconnect_updateFlowSourceCmd.Flags().String("max-latency", "", "The maximum latency in milliseconds.")
		mediaconnect_updateFlowSourceCmd.Flags().String("max-sync-buffer", "", "The size of the buffer (in milliseconds) to use to sync incoming source data.")
		mediaconnect_updateFlowSourceCmd.Flags().String("media-stream-source-configurations", "", "The media stream that is associated with the source, and the parameters for that association.")
		mediaconnect_updateFlowSourceCmd.Flags().String("min-latency", "", "The minimum latency in milliseconds for SRT-based streams.")
		mediaconnect_updateFlowSourceCmd.Flags().String("protocol", "", "The protocol that the source uses to deliver the content to MediaConnect.")
		mediaconnect_updateFlowSourceCmd.Flags().String("sender-control-port", "", "The port that the flow uses to send outbound requests to initiate connection with the sender.")
		mediaconnect_updateFlowSourceCmd.Flags().String("sender-ip-address", "", "The IP address that the flow communicates with to initiate connection with the sender.")
		mediaconnect_updateFlowSourceCmd.Flags().String("source-arn", "", "The ARN of the source that you want to update.")
		mediaconnect_updateFlowSourceCmd.Flags().String("source-listener-address", "", "The source IP or domain name for SRT-caller protocol.")
		mediaconnect_updateFlowSourceCmd.Flags().String("source-listener-port", "", "Source port for SRT-caller protocol.")
		mediaconnect_updateFlowSourceCmd.Flags().String("stream-id", "", "The stream ID that you want to use for this transport.")
		mediaconnect_updateFlowSourceCmd.Flags().String("vpc-interface-name", "", "The name of the VPC interface that you want to send your output to.")
		mediaconnect_updateFlowSourceCmd.Flags().String("whitelist-cidr", "", "The range of IP addresses that are allowed to contribute content to your source.")
		mediaconnect_updateFlowSourceCmd.MarkFlagRequired("flow-arn")
		mediaconnect_updateFlowSourceCmd.MarkFlagRequired("source-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_updateFlowSourceCmd)
}
