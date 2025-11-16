package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_listExecutionsCmd = &cobra.Command{
	Use:   "list-executions",
	Short: "Lists all executions of a state machine or a Map Run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_listExecutionsCmd).Standalone()

	stepfunctions_listExecutionsCmd.Flags().String("map-run-arn", "", "The Amazon Resource Name (ARN) of the Map Run that started the child workflow executions.")
	stepfunctions_listExecutionsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	stepfunctions_listExecutionsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	stepfunctions_listExecutionsCmd.Flags().String("redrive-filter", "", "Sets a filter to list executions based on whether or not they have been redriven.")
	stepfunctions_listExecutionsCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine whose executions is listed.")
	stepfunctions_listExecutionsCmd.Flags().String("status-filter", "", "If specified, only list the executions whose current execution status matches the given filter.")
	stepfunctionsCmd.AddCommand(stepfunctions_listExecutionsCmd)
}
