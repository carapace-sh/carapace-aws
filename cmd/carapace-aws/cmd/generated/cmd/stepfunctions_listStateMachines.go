package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_listStateMachinesCmd = &cobra.Command{
	Use:   "list-state-machines",
	Short: "Lists the existing state machines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_listStateMachinesCmd).Standalone()

	stepfunctions_listStateMachinesCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	stepfunctions_listStateMachinesCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	stepfunctionsCmd.AddCommand(stepfunctions_listStateMachinesCmd)
}
