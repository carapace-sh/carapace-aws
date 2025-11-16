package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_getExecutionFlowSnapshotCmd = &cobra.Command{
	Use:   "get-execution-flow-snapshot",
	Short: "Retrieves the flow definition snapshot used for a flow execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_getExecutionFlowSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_getExecutionFlowSnapshotCmd).Standalone()

		bedrockAgentRuntime_getExecutionFlowSnapshotCmd.Flags().String("execution-identifier", "", "The unique identifier of the flow execution.")
		bedrockAgentRuntime_getExecutionFlowSnapshotCmd.Flags().String("flow-alias-identifier", "", "The unique identifier of the flow alias used for the flow execution.")
		bedrockAgentRuntime_getExecutionFlowSnapshotCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgentRuntime_getExecutionFlowSnapshotCmd.MarkFlagRequired("execution-identifier")
		bedrockAgentRuntime_getExecutionFlowSnapshotCmd.MarkFlagRequired("flow-alias-identifier")
		bedrockAgentRuntime_getExecutionFlowSnapshotCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_getExecutionFlowSnapshotCmd)
}
