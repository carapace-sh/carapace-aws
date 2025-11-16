package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerA2iRuntime_deleteHumanLoopCmd = &cobra.Command{
	Use:   "delete-human-loop",
	Short: "Deletes the specified human loop for a flow definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerA2iRuntime_deleteHumanLoopCmd).Standalone()

	sagemakerA2iRuntime_deleteHumanLoopCmd.Flags().String("human-loop-name", "", "The name of the human loop that you want to delete.")
	sagemakerA2iRuntime_deleteHumanLoopCmd.MarkFlagRequired("human-loop-name")
	sagemakerA2iRuntimeCmd.AddCommand(sagemakerA2iRuntime_deleteHumanLoopCmd)
}
