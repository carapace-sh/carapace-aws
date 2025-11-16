package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerA2iRuntime_listHumanLoopsCmd = &cobra.Command{
	Use:   "list-human-loops",
	Short: "Returns information about human loops, given the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerA2iRuntime_listHumanLoopsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerA2iRuntime_listHumanLoopsCmd).Standalone()

		sagemakerA2iRuntime_listHumanLoopsCmd.Flags().String("creation-time-after", "", "(Optional) The timestamp of the date when you want the human loops to begin in ISO 8601 format.")
		sagemakerA2iRuntime_listHumanLoopsCmd.Flags().String("creation-time-before", "", "(Optional) The timestamp of the date before which you want the human loops to begin in ISO 8601 format.")
		sagemakerA2iRuntime_listHumanLoopsCmd.Flags().String("flow-definition-arn", "", "The Amazon Resource Name (ARN) of a flow definition.")
		sagemakerA2iRuntime_listHumanLoopsCmd.Flags().String("max-results", "", "The total number of items to return.")
		sagemakerA2iRuntime_listHumanLoopsCmd.Flags().String("next-token", "", "A token to display the next page of results.")
		sagemakerA2iRuntime_listHumanLoopsCmd.Flags().String("sort-order", "", "Optional.")
		sagemakerA2iRuntime_listHumanLoopsCmd.MarkFlagRequired("flow-definition-arn")
	})
	sagemakerA2iRuntimeCmd.AddCommand(sagemakerA2iRuntime_listHumanLoopsCmd)
}
