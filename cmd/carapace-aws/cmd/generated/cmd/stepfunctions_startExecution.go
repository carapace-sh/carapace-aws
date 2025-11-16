package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_startExecutionCmd = &cobra.Command{
	Use:   "start-execution",
	Short: "Starts a state machine execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_startExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_startExecutionCmd).Standalone()

		stepfunctions_startExecutionCmd.Flags().String("input", "", "The string that contains the JSON input data for the execution, for example:")
		stepfunctions_startExecutionCmd.Flags().String("name", "", "Optional name of the execution.")
		stepfunctions_startExecutionCmd.Flags().String("state-machine-arn", "", "The Amazon Resource Name (ARN) of the state machine to execute.")
		stepfunctions_startExecutionCmd.Flags().String("trace-header", "", "Passes the X-Ray trace header.")
		stepfunctions_startExecutionCmd.MarkFlagRequired("state-machine-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_startExecutionCmd)
}
