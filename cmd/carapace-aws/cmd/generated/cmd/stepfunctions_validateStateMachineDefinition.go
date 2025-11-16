package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_validateStateMachineDefinitionCmd = &cobra.Command{
	Use:   "validate-state-machine-definition",
	Short: "Validates the syntax of a state machine definition specified in [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) (ASL), a JSON-based, structured language.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_validateStateMachineDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_validateStateMachineDefinitionCmd).Standalone()

		stepfunctions_validateStateMachineDefinitionCmd.Flags().String("definition", "", "The Amazon States Language definition of the state machine.")
		stepfunctions_validateStateMachineDefinitionCmd.Flags().String("max-results", "", "The maximum number of diagnostics that are returned per call.")
		stepfunctions_validateStateMachineDefinitionCmd.Flags().String("severity", "", "Minimum level of diagnostics to return.")
		stepfunctions_validateStateMachineDefinitionCmd.Flags().String("type", "", "The target type of state machine for this definition.")
		stepfunctions_validateStateMachineDefinitionCmd.MarkFlagRequired("definition")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_validateStateMachineDefinitionCmd)
}
