package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_listMapRunsCmd = &cobra.Command{
	Use:   "list-map-runs",
	Short: "Lists all Map Runs that were started by a given state machine execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_listMapRunsCmd).Standalone()

	stepfunctions_listMapRunsCmd.Flags().String("execution-arn", "", "The Amazon Resource Name (ARN) of the execution for which the Map Runs must be listed.")
	stepfunctions_listMapRunsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	stepfunctions_listMapRunsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	stepfunctions_listMapRunsCmd.MarkFlagRequired("execution-arn")
	stepfunctionsCmd.AddCommand(stepfunctions_listMapRunsCmd)
}
