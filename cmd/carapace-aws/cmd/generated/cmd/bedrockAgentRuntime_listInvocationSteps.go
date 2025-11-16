package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_listInvocationStepsCmd = &cobra.Command{
	Use:   "list-invocation-steps",
	Short: "Lists all invocation steps associated with a session and optionally, an invocation within the session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_listInvocationStepsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_listInvocationStepsCmd).Standalone()

		bedrockAgentRuntime_listInvocationStepsCmd.Flags().String("invocation-identifier", "", "The unique identifier (in UUID format) for the invocation to list invocation steps for.")
		bedrockAgentRuntime_listInvocationStepsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgentRuntime_listInvocationStepsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgentRuntime_listInvocationStepsCmd.Flags().String("session-identifier", "", "The unique identifier for the session associated with the invocation steps.")
		bedrockAgentRuntime_listInvocationStepsCmd.MarkFlagRequired("session-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_listInvocationStepsCmd)
}
