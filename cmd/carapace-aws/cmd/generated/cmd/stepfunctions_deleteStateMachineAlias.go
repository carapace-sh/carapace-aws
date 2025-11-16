package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_deleteStateMachineAliasCmd = &cobra.Command{
	Use:   "delete-state-machine-alias",
	Short: "Deletes a state machine [alias](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html).\n\nAfter you delete a state machine alias, you can't use it to start executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_deleteStateMachineAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_deleteStateMachineAliasCmd).Standalone()

		stepfunctions_deleteStateMachineAliasCmd.Flags().String("state-machine-alias-arn", "", "The Amazon Resource Name (ARN) of the state machine alias to delete.")
		stepfunctions_deleteStateMachineAliasCmd.MarkFlagRequired("state-machine-alias-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_deleteStateMachineAliasCmd)
}
