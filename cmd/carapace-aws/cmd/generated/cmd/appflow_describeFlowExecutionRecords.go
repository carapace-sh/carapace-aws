package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_describeFlowExecutionRecordsCmd = &cobra.Command{
	Use:   "describe-flow-execution-records",
	Short: "Fetches the execution history of the flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_describeFlowExecutionRecordsCmd).Standalone()

	appflow_describeFlowExecutionRecordsCmd.Flags().String("flow-name", "", "The specified name of the flow.")
	appflow_describeFlowExecutionRecordsCmd.Flags().String("max-results", "", "Specifies the maximum number of items that should be returned in the result set.")
	appflow_describeFlowExecutionRecordsCmd.Flags().String("next-token", "", "The pagination token for the next page of data.")
	appflow_describeFlowExecutionRecordsCmd.MarkFlagRequired("flow-name")
	appflowCmd.AddCommand(appflow_describeFlowExecutionRecordsCmd)
}
