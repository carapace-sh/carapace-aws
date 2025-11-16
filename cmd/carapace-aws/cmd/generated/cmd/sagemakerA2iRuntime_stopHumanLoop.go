package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerA2iRuntime_stopHumanLoopCmd = &cobra.Command{
	Use:   "stop-human-loop",
	Short: "Stops the specified human loop.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerA2iRuntime_stopHumanLoopCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerA2iRuntime_stopHumanLoopCmd).Standalone()

		sagemakerA2iRuntime_stopHumanLoopCmd.Flags().String("human-loop-name", "", "The name of the human loop that you want to stop.")
		sagemakerA2iRuntime_stopHumanLoopCmd.MarkFlagRequired("human-loop-name")
	})
	sagemakerA2iRuntimeCmd.AddCommand(sagemakerA2iRuntime_stopHumanLoopCmd)
}
