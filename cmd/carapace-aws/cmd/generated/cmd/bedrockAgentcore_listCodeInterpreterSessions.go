package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_listCodeInterpreterSessionsCmd = &cobra.Command{
	Use:   "list-code-interpreter-sessions",
	Short: "Retrieves a list of code interpreter sessions in Amazon Bedrock that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_listCodeInterpreterSessionsCmd).Standalone()

	bedrockAgentcore_listCodeInterpreterSessionsCmd.Flags().String("code-interpreter-identifier", "", "The unique identifier of the code interpreter to list sessions for.")
	bedrockAgentcore_listCodeInterpreterSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	bedrockAgentcore_listCodeInterpreterSessionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	bedrockAgentcore_listCodeInterpreterSessionsCmd.Flags().String("status", "", "The status of the code interpreter sessions to list.")
	bedrockAgentcore_listCodeInterpreterSessionsCmd.MarkFlagRequired("code-interpreter-identifier")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_listCodeInterpreterSessionsCmd)
}
