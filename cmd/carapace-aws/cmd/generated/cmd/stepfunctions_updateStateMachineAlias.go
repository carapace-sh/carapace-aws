package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_updateStateMachineAliasCmd = &cobra.Command{
	Use:   "update-state-machine-alias",
	Short: "Updates the configuration of an existing state machine [alias](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html) by modifying its `description` or `routingConfiguration`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_updateStateMachineAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_updateStateMachineAliasCmd).Standalone()

		stepfunctions_updateStateMachineAliasCmd.Flags().String("description", "", "A description of the state machine alias.")
		stepfunctions_updateStateMachineAliasCmd.Flags().String("routing-configuration", "", "The routing configuration of the state machine alias.")
		stepfunctions_updateStateMachineAliasCmd.Flags().String("state-machine-alias-arn", "", "The Amazon Resource Name (ARN) of the state machine alias.")
		stepfunctions_updateStateMachineAliasCmd.MarkFlagRequired("state-machine-alias-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_updateStateMachineAliasCmd)
}
