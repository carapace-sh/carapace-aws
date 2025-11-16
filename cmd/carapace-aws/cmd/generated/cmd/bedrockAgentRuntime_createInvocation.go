package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_createInvocationCmd = &cobra.Command{
	Use:   "create-invocation",
	Short: "Creates a new invocation within a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_createInvocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_createInvocationCmd).Standalone()

		bedrockAgentRuntime_createInvocationCmd.Flags().String("description", "", "A description for the interactions in the invocation.")
		bedrockAgentRuntime_createInvocationCmd.Flags().String("invocation-id", "", "A unique identifier for the invocation in UUID format.")
		bedrockAgentRuntime_createInvocationCmd.Flags().String("session-identifier", "", "The unique identifier for the associated session for the invocation.")
		bedrockAgentRuntime_createInvocationCmd.MarkFlagRequired("session-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_createInvocationCmd)
}
