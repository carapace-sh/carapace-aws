package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeExecutionCmd = &cobra.Command{
	Use:   "describe-execution",
	Short: "Retrieves information about the execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeExecutionCmd).Standalone()

	iotsitewise_describeExecutionCmd.Flags().String("execution-id", "", "The ID of the execution.")
	iotsitewise_describeExecutionCmd.MarkFlagRequired("execution-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeExecutionCmd)
}
