package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_deleteStateMachineVersionCmd = &cobra.Command{
	Use:   "delete-state-machine-version",
	Short: "Deletes a state machine [version](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html). After you delete a version, you can't call [StartExecution]() using that version's ARN or use the version with a state machine [alias](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_deleteStateMachineVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_deleteStateMachineVersionCmd).Standalone()

		stepfunctions_deleteStateMachineVersionCmd.Flags().String("state-machine-version-arn", "", "The Amazon Resource Name (ARN) of the state machine version to delete.")
		stepfunctions_deleteStateMachineVersionCmd.MarkFlagRequired("state-machine-version-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_deleteStateMachineVersionCmd)
}
