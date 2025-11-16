package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_searchFlowExecutionsCmd = &cobra.Command{
	Use:   "search-flow-executions",
	Short: "Searches for AWS IoT Things Graph workflow execution instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_searchFlowExecutionsCmd).Standalone()

	iotthingsgraph_searchFlowExecutionsCmd.Flags().String("end-time", "", "The date and time of the latest flow execution to return.")
	iotthingsgraph_searchFlowExecutionsCmd.Flags().String("flow-execution-id", "", "The ID of a flow execution.")
	iotthingsgraph_searchFlowExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	iotthingsgraph_searchFlowExecutionsCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iotthingsgraph_searchFlowExecutionsCmd.Flags().String("start-time", "", "The date and time of the earliest flow execution to return.")
	iotthingsgraph_searchFlowExecutionsCmd.Flags().String("system-instance-id", "", "The ID of the system instance that contains the flow.")
	iotthingsgraph_searchFlowExecutionsCmd.MarkFlagRequired("system-instance-id")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_searchFlowExecutionsCmd)
}
