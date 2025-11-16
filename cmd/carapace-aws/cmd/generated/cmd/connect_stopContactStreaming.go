package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_stopContactStreamingCmd = &cobra.Command{
	Use:   "stop-contact-streaming",
	Short: "Ends message streaming on a specified contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_stopContactStreamingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_stopContactStreamingCmd).Standalone()

		connect_stopContactStreamingCmd.Flags().String("contact-id", "", "The identifier of the contact.")
		connect_stopContactStreamingCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_stopContactStreamingCmd.Flags().String("streaming-id", "", "The identifier of the streaming configuration enabled.")
		connect_stopContactStreamingCmd.MarkFlagRequired("contact-id")
		connect_stopContactStreamingCmd.MarkFlagRequired("instance-id")
		connect_stopContactStreamingCmd.MarkFlagRequired("streaming-id")
	})
	connectCmd.AddCommand(connect_stopContactStreamingCmd)
}
