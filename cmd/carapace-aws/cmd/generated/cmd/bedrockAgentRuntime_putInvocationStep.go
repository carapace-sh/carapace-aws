package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_putInvocationStepCmd = &cobra.Command{
	Use:   "put-invocation-step",
	Short: "Add an invocation step to an invocation in a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_putInvocationStepCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_putInvocationStepCmd).Standalone()

		bedrockAgentRuntime_putInvocationStepCmd.Flags().String("invocation-identifier", "", "The unique identifier (in UUID format) of the invocation to add the invocation step to.")
		bedrockAgentRuntime_putInvocationStepCmd.Flags().String("invocation-step-id", "", "The unique identifier of the invocation step in UUID format.")
		bedrockAgentRuntime_putInvocationStepCmd.Flags().String("invocation-step-time", "", "The timestamp for when the invocation step occurred.")
		bedrockAgentRuntime_putInvocationStepCmd.Flags().String("payload", "", "The payload for the invocation step, including text and images for the interaction.")
		bedrockAgentRuntime_putInvocationStepCmd.Flags().String("session-identifier", "", "The unique identifier for the session to add the invocation step to.")
		bedrockAgentRuntime_putInvocationStepCmd.MarkFlagRequired("invocation-identifier")
		bedrockAgentRuntime_putInvocationStepCmd.MarkFlagRequired("invocation-step-time")
		bedrockAgentRuntime_putInvocationStepCmd.MarkFlagRequired("payload")
		bedrockAgentRuntime_putInvocationStepCmd.MarkFlagRequired("session-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_putInvocationStepCmd)
}
