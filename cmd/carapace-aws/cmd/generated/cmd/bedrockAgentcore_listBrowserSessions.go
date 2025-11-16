package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_listBrowserSessionsCmd = &cobra.Command{
	Use:   "list-browser-sessions",
	Short: "Retrieves a list of browser sessions in Amazon Bedrock that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_listBrowserSessionsCmd).Standalone()

	bedrockAgentcore_listBrowserSessionsCmd.Flags().String("browser-identifier", "", "The unique identifier of the browser to list sessions for.")
	bedrockAgentcore_listBrowserSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	bedrockAgentcore_listBrowserSessionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	bedrockAgentcore_listBrowserSessionsCmd.Flags().String("status", "", "The status of the browser sessions to list.")
	bedrockAgentcore_listBrowserSessionsCmd.MarkFlagRequired("browser-identifier")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_listBrowserSessionsCmd)
}
