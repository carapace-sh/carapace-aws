package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerA2iRuntime_describeHumanLoopCmd = &cobra.Command{
	Use:   "describe-human-loop",
	Short: "Returns information about the specified human loop.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerA2iRuntime_describeHumanLoopCmd).Standalone()

	sagemakerA2iRuntime_describeHumanLoopCmd.Flags().String("human-loop-name", "", "The name of the human loop that you want information about.")
	sagemakerA2iRuntime_describeHumanLoopCmd.MarkFlagRequired("human-loop-name")
	sagemakerA2iRuntimeCmd.AddCommand(sagemakerA2iRuntime_describeHumanLoopCmd)
}
