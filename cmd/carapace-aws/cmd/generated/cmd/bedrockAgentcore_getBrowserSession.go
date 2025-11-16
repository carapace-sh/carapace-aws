package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getBrowserSessionCmd = &cobra.Command{
	Use:   "get-browser-session",
	Short: "Retrieves detailed information about a specific browser session in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getBrowserSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_getBrowserSessionCmd).Standalone()

		bedrockAgentcore_getBrowserSessionCmd.Flags().String("browser-identifier", "", "The unique identifier of the browser associated with the session.")
		bedrockAgentcore_getBrowserSessionCmd.Flags().String("session-id", "", "The unique identifier of the browser session to retrieve.")
		bedrockAgentcore_getBrowserSessionCmd.MarkFlagRequired("browser-identifier")
		bedrockAgentcore_getBrowserSessionCmd.MarkFlagRequired("session-id")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getBrowserSessionCmd)
}
