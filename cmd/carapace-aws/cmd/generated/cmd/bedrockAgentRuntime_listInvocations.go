package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_listInvocationsCmd = &cobra.Command{
	Use:   "list-invocations",
	Short: "Lists all invocations associated with a specific session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_listInvocationsCmd).Standalone()

	bedrockAgentRuntime_listInvocationsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgentRuntime_listInvocationsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrockAgentRuntime_listInvocationsCmd.Flags().String("session-identifier", "", "The unique identifier for the session to list invocations for.")
	bedrockAgentRuntime_listInvocationsCmd.MarkFlagRequired("session-identifier")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_listInvocationsCmd)
}
