package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_testStateCmd = &cobra.Command{
	Use:   "test-state",
	Short: "Accepts the definition of a single state and executes it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_testStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_testStateCmd).Standalone()

		stepfunctions_testStateCmd.Flags().String("definition", "", "The [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) (ASL) definition of the state.")
		stepfunctions_testStateCmd.Flags().String("input", "", "A string that contains the JSON input data for the state.")
		stepfunctions_testStateCmd.Flags().String("inspection-level", "", "Determines the values to return when a state is tested.")
		stepfunctions_testStateCmd.Flags().String("reveal-secrets", "", "Specifies whether or not to include secret information in the test result.")
		stepfunctions_testStateCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the execution role with the required IAM permissions for the state.")
		stepfunctions_testStateCmd.Flags().String("variables", "", "JSON object literal that sets variables used in the state under test.")
		stepfunctions_testStateCmd.MarkFlagRequired("definition")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_testStateCmd)
}
