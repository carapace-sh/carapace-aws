package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_createStateMachineCmd = &cobra.Command{
	Use:   "create-state-machine",
	Short: "Creates a state machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_createStateMachineCmd).Standalone()

	stepfunctions_createStateMachineCmd.Flags().String("definition", "", "The Amazon States Language definition of the state machine.")
	stepfunctions_createStateMachineCmd.Flags().String("encryption-configuration", "", "Settings to configure server-side encryption.")
	stepfunctions_createStateMachineCmd.Flags().String("logging-configuration", "", "Defines what execution history events are logged and where they are logged.")
	stepfunctions_createStateMachineCmd.Flags().String("name", "", "The name of the state machine.")
	stepfunctions_createStateMachineCmd.Flags().String("publish", "", "Set to `true` to publish the first version of the state machine during creation.")
	stepfunctions_createStateMachineCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to use for this state machine.")
	stepfunctions_createStateMachineCmd.Flags().String("tags", "", "Tags to be added when creating a state machine.")
	stepfunctions_createStateMachineCmd.Flags().String("tracing-configuration", "", "Selects whether X-Ray tracing is enabled.")
	stepfunctions_createStateMachineCmd.Flags().String("type", "", "Determines whether a Standard or Express state machine is created.")
	stepfunctions_createStateMachineCmd.Flags().String("version-description", "", "Sets description about the state machine version.")
	stepfunctions_createStateMachineCmd.MarkFlagRequired("definition")
	stepfunctions_createStateMachineCmd.MarkFlagRequired("name")
	stepfunctions_createStateMachineCmd.MarkFlagRequired("role-arn")
	stepfunctionsCmd.AddCommand(stepfunctions_createStateMachineCmd)
}
