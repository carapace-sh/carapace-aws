package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startContactStreamingCmd = &cobra.Command{
	Use:   "start-contact-streaming",
	Short: "Initiates real-time message streaming for a new chat contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startContactStreamingCmd).Standalone()

	connect_startContactStreamingCmd.Flags().String("chat-streaming-configuration", "", "The streaming configuration, such as the Amazon SNS streaming endpoint.")
	connect_startContactStreamingCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_startContactStreamingCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_startContactStreamingCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_startContactStreamingCmd.MarkFlagRequired("chat-streaming-configuration")
	connect_startContactStreamingCmd.MarkFlagRequired("client-token")
	connect_startContactStreamingCmd.MarkFlagRequired("contact-id")
	connect_startContactStreamingCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_startContactStreamingCmd)
}
