package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_getInvocationStepCmd = &cobra.Command{
	Use:   "get-invocation-step",
	Short: "Retrieves the details of a specific invocation step within an invocation in a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_getInvocationStepCmd).Standalone()

	bedrockAgentRuntime_getInvocationStepCmd.Flags().String("invocation-identifier", "", "The unique identifier for the invocation in UUID format.")
	bedrockAgentRuntime_getInvocationStepCmd.Flags().String("invocation-step-id", "", "The unique identifier (in UUID format) for the specific invocation step to retrieve.")
	bedrockAgentRuntime_getInvocationStepCmd.Flags().String("session-identifier", "", "The unique identifier for the invocation step's associated session.")
	bedrockAgentRuntime_getInvocationStepCmd.MarkFlagRequired("invocation-identifier")
	bedrockAgentRuntime_getInvocationStepCmd.MarkFlagRequired("invocation-step-id")
	bedrockAgentRuntime_getInvocationStepCmd.MarkFlagRequired("session-identifier")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_getInvocationStepCmd)
}
