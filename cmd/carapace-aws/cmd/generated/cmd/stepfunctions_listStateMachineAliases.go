package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_listStateMachineAliasesCmd = &cobra.Command{
	Use:   "list-state-machine-aliases",
	Short: "Lists [aliases](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html) for a specified state machine ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_listStateMachineAliasesCmd).Standalone()

	stepfunctions_listStateMachineAliasesCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	stepfunctions_listStateMachineAliasesCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	stepfunctions_listStateMachineAliasesCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine for which you want to list aliases.")
	stepfunctions_listStateMachineAliasesCmd.MarkFlagRequired("state-machine-arn")
	stepfunctionsCmd.AddCommand(stepfunctions_listStateMachineAliasesCmd)
}
