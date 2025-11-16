package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerA2iRuntime_startHumanLoopCmd = &cobra.Command{
	Use:   "start-human-loop",
	Short: "Starts a human loop, provided that at least one activation condition is met.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerA2iRuntime_startHumanLoopCmd).Standalone()

	sagemakerA2iRuntime_startHumanLoopCmd.Flags().String("data-attributes", "", "Attributes of the specified data.")
	sagemakerA2iRuntime_startHumanLoopCmd.Flags().String("flow-definition-arn", "", "The Amazon Resource Name (ARN) of the flow definition associated with this human loop.")
	sagemakerA2iRuntime_startHumanLoopCmd.Flags().String("human-loop-input", "", "An object that contains information about the human loop.")
	sagemakerA2iRuntime_startHumanLoopCmd.Flags().String("human-loop-name", "", "The name of the human loop.")
	sagemakerA2iRuntime_startHumanLoopCmd.MarkFlagRequired("flow-definition-arn")
	sagemakerA2iRuntime_startHumanLoopCmd.MarkFlagRequired("human-loop-input")
	sagemakerA2iRuntime_startHumanLoopCmd.MarkFlagRequired("human-loop-name")
	sagemakerA2iRuntimeCmd.AddCommand(sagemakerA2iRuntime_startHumanLoopCmd)
}
