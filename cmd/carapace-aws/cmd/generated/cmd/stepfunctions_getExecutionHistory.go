package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_getExecutionHistoryCmd = &cobra.Command{
	Use:   "get-execution-history",
	Short: "Returns the history of the specified execution as a list of events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_getExecutionHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_getExecutionHistoryCmd).Standalone()

		stepfunctions_getExecutionHistoryCmd.Flags().String("execution-arn", "", "The Amazon Resource Name (ARN) of the execution.")
		stepfunctions_getExecutionHistoryCmd.Flags().String("include-execution-data", "", "You can select whether execution data (input or output of a history event) is returned.")
		stepfunctions_getExecutionHistoryCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		stepfunctions_getExecutionHistoryCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
		stepfunctions_getExecutionHistoryCmd.Flags().String("reverse-order", "", "Lists events in descending order of their `timeStamp`.")
		stepfunctions_getExecutionHistoryCmd.MarkFlagRequired("execution-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_getExecutionHistoryCmd)
}
