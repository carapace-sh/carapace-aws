package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createStreamingUrlCmd = &cobra.Command{
	Use:   "create-streaming-url",
	Short: "Creates a temporary URL to start an AppStream 2.0 streaming session for the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createStreamingUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createStreamingUrlCmd).Standalone()

		appstream_createStreamingUrlCmd.Flags().String("application-id", "", "The name of the application to launch after the session starts.")
		appstream_createStreamingUrlCmd.Flags().String("fleet-name", "", "The name of the fleet.")
		appstream_createStreamingUrlCmd.Flags().String("session-context", "", "The session context.")
		appstream_createStreamingUrlCmd.Flags().String("stack-name", "", "The name of the stack.")
		appstream_createStreamingUrlCmd.Flags().String("user-id", "", "The identifier of the user.")
		appstream_createStreamingUrlCmd.Flags().String("validity", "", "The time that the streaming URL will be valid, in seconds.")
		appstream_createStreamingUrlCmd.MarkFlagRequired("fleet-name")
		appstream_createStreamingUrlCmd.MarkFlagRequired("stack-name")
		appstream_createStreamingUrlCmd.MarkFlagRequired("user-id")
	})
	appstreamCmd.AddCommand(appstream_createStreamingUrlCmd)
}
