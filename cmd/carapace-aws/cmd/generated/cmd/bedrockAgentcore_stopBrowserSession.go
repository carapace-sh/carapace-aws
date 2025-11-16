package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_stopBrowserSessionCmd = &cobra.Command{
	Use:   "stop-browser-session",
	Short: "Terminates an active browser session in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_stopBrowserSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_stopBrowserSessionCmd).Standalone()

		bedrockAgentcore_stopBrowserSessionCmd.Flags().String("browser-identifier", "", "The unique identifier of the browser associated with the session.")
		bedrockAgentcore_stopBrowserSessionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgentcore_stopBrowserSessionCmd.Flags().String("session-id", "", "The unique identifier of the browser session to stop.")
		bedrockAgentcore_stopBrowserSessionCmd.Flags().String("trace-id", "", "The trace identifier for request tracking.")
		bedrockAgentcore_stopBrowserSessionCmd.Flags().String("trace-parent", "", "The parent trace information for distributed tracing.")
		bedrockAgentcore_stopBrowserSessionCmd.MarkFlagRequired("browser-identifier")
		bedrockAgentcore_stopBrowserSessionCmd.MarkFlagRequired("session-id")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_stopBrowserSessionCmd)
}
