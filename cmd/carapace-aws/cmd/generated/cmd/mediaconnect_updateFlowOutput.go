package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateFlowOutputCmd = &cobra.Command{
	Use:   "update-flow-output",
	Short: "Updates an existing flow output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateFlowOutputCmd).Standalone()

	mediaconnect_updateFlowOutputCmd.Flags().String("cidr-allow-list", "", "The range of IP addresses that should be allowed to initiate output requests to this flow.")
	mediaconnect_updateFlowOutputCmd.Flags().String("description", "", "A description of the output.")
	mediaconnect_updateFlowOutputCmd.Flags().String("destination", "", "The IP address where you want to send the output.")
	mediaconnect_updateFlowOutputCmd.Flags().String("encryption", "", "The type of key used for the encryption.")
	mediaconnect_updateFlowOutputCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that is associated with the output that you want to update.")
	mediaconnect_updateFlowOutputCmd.Flags().String("max-latency", "", "The maximum latency in milliseconds.")
	mediaconnect_updateFlowOutputCmd.Flags().String("media-stream-output-configurations", "", "The media streams that are associated with the output, and the parameters for those associations.")
	mediaconnect_updateFlowOutputCmd.Flags().String("min-latency", "", "The minimum latency in milliseconds for SRT-based streams.")
	mediaconnect_updateFlowOutputCmd.Flags().String("ndi-program-name", "", "A suffix for the names of the NDI sources that the flow creates.")
	mediaconnect_updateFlowOutputCmd.Flags().String("ndi-speed-hq-quality", "", "A quality setting for the NDI Speed HQ encoder.")
	mediaconnect_updateFlowOutputCmd.Flags().String("output-arn", "", "The ARN of the output that you want to update.")
	mediaconnect_updateFlowOutputCmd.Flags().String("output-status", "", "An indication of whether the output should transmit data or not.")
	mediaconnect_updateFlowOutputCmd.Flags().String("port", "", "The port to use when content is distributed to this output.")
	mediaconnect_updateFlowOutputCmd.Flags().String("protocol", "", "The protocol to use for the output.")
	mediaconnect_updateFlowOutputCmd.Flags().String("remote-id", "", "The remote ID for the Zixi-pull stream.")
	mediaconnect_updateFlowOutputCmd.Flags().String("sender-control-port", "", "The port that the flow uses to send outbound requests to initiate connection with the sender.")
	mediaconnect_updateFlowOutputCmd.Flags().String("sender-ip-address", "", "The IP address that the flow communicates with to initiate connection with the sender.")
	mediaconnect_updateFlowOutputCmd.Flags().String("smoothing-latency", "", "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.")
	mediaconnect_updateFlowOutputCmd.Flags().String("stream-id", "", "The stream ID that you want to use for this transport.")
	mediaconnect_updateFlowOutputCmd.Flags().String("vpc-interface-attachment", "", "The name of the VPC interface attachment to use for this output.")
	mediaconnect_updateFlowOutputCmd.MarkFlagRequired("flow-arn")
	mediaconnect_updateFlowOutputCmd.MarkFlagRequired("output-arn")
	mediaconnectCmd.AddCommand(mediaconnect_updateFlowOutputCmd)
}
