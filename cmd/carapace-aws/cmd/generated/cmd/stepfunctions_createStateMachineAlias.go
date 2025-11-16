package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_createStateMachineAliasCmd = &cobra.Command{
	Use:   "create-state-machine-alias",
	Short: "Creates an [alias](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html) for a state machine that points to one or two [versions](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html) of the same state machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_createStateMachineAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_createStateMachineAliasCmd).Standalone()

		stepfunctions_createStateMachineAliasCmd.Flags().String("description", "", "A description for the state machine alias.")
		stepfunctions_createStateMachineAliasCmd.Flags().String("name", "", "The name of the state machine alias.")
		stepfunctions_createStateMachineAliasCmd.Flags().String("routing-configuration", "", "The routing configuration of a state machine alias.")
		stepfunctions_createStateMachineAliasCmd.MarkFlagRequired("name")
		stepfunctions_createStateMachineAliasCmd.MarkFlagRequired("routing-configuration")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_createStateMachineAliasCmd)
}
