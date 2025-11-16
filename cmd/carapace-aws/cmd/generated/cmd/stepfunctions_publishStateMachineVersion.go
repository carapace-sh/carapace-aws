package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_publishStateMachineVersionCmd = &cobra.Command{
	Use:   "publish-state-machine-version",
	Short: "Creates a [version](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html) from the current revision of a state machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_publishStateMachineVersionCmd).Standalone()

	stepfunctions_publishStateMachineVersionCmd.Flags().String("description", "", "An optional description of the state machine version.")
	stepfunctions_publishStateMachineVersionCmd.Flags().String("revision-id", "", "Only publish the state machine version if the current state machine's revision ID matches the specified ID.")
	stepfunctions_publishStateMachineVersionCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine.")
	stepfunctions_publishStateMachineVersionCmd.MarkFlagRequired("state-machine-arn")
	stepfunctionsCmd.AddCommand(stepfunctions_publishStateMachineVersionCmd)
}
