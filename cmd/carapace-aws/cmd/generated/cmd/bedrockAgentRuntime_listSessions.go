package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "Lists all sessions in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_listSessionsCmd).Standalone()

	bedrockAgentRuntime_listSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgentRuntime_listSessionsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_listSessionsCmd)
}
