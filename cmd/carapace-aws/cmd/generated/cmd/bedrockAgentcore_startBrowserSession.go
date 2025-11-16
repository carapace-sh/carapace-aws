package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_startBrowserSessionCmd = &cobra.Command{
	Use:   "start-browser-session",
	Short: "Creates and initializes a browser session in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_startBrowserSessionCmd).Standalone()

	bedrockAgentcore_startBrowserSessionCmd.Flags().String("browser-identifier", "", "The unique identifier of the browser to use for this session.")
	bedrockAgentcore_startBrowserSessionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgentcore_startBrowserSessionCmd.Flags().String("name", "", "The name of the browser session.")
	bedrockAgentcore_startBrowserSessionCmd.Flags().String("session-timeout-seconds", "", "The time in seconds after which the session automatically terminates if there is no activity.")
	bedrockAgentcore_startBrowserSessionCmd.Flags().String("trace-id", "", "The trace identifier for request tracking.")
	bedrockAgentcore_startBrowserSessionCmd.Flags().String("trace-parent", "", "The parent trace information for distributed tracing.")
	bedrockAgentcore_startBrowserSessionCmd.Flags().String("view-port", "", "The dimensions of the browser viewport for this session.")
	bedrockAgentcore_startBrowserSessionCmd.MarkFlagRequired("browser-identifier")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_startBrowserSessionCmd)
}
