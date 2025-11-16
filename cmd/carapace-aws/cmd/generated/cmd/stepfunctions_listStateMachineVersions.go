package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_listStateMachineVersionsCmd = &cobra.Command{
	Use:   "list-state-machine-versions",
	Short: "Lists [versions](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html) for the specified state machine Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_listStateMachineVersionsCmd).Standalone()

	stepfunctions_listStateMachineVersionsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	stepfunctions_listStateMachineVersionsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	stepfunctions_listStateMachineVersionsCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine.")
	stepfunctions_listStateMachineVersionsCmd.MarkFlagRequired("state-machine-arn")
	stepfunctionsCmd.AddCommand(stepfunctions_listStateMachineVersionsCmd)
}
