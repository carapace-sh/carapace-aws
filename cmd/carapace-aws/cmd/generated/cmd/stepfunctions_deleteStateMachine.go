package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_deleteStateMachineCmd = &cobra.Command{
	Use:   "delete-state-machine",
	Short: "Deletes a state machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_deleteStateMachineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_deleteStateMachineCmd).Standalone()

		stepfunctions_deleteStateMachineCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine to delete.")
		stepfunctions_deleteStateMachineCmd.MarkFlagRequired("state-machine-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_deleteStateMachineCmd)
}
