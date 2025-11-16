package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_describeStateMachineAliasCmd = &cobra.Command{
	Use:   "describe-state-machine-alias",
	Short: "Returns details about a state machine [alias](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html).\n\n**Related operations:**\n\n- [CreateStateMachineAlias]()\n- [ListStateMachineAliases]()\n- [UpdateStateMachineAlias]()\n- [DeleteStateMachineAlias]()",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_describeStateMachineAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_describeStateMachineAliasCmd).Standalone()

		stepfunctions_describeStateMachineAliasCmd.Flags().String("state-machine-alias-arn", "", "The Amazon Resource Name (ARN) of the state machine alias.")
		stepfunctions_describeStateMachineAliasCmd.MarkFlagRequired("state-machine-alias-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_describeStateMachineAliasCmd)
}
