package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_listFlowExecutionMessagesCmd = &cobra.Command{
	Use:   "list-flow-execution-messages",
	Short: "Returns a list of objects that contain information about events in a flow execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_listFlowExecutionMessagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_listFlowExecutionMessagesCmd).Standalone()

		iotthingsgraph_listFlowExecutionMessagesCmd.Flags().String("flow-execution-id", "", "The ID of the flow execution.")
		iotthingsgraph_listFlowExecutionMessagesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		iotthingsgraph_listFlowExecutionMessagesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iotthingsgraph_listFlowExecutionMessagesCmd.MarkFlagRequired("flow-execution-id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_listFlowExecutionMessagesCmd)
}
