package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeExecutionCmd = &cobra.Command{
	Use:   "describe-execution",
	Short: "You can use `DescribeExecution` to check the details of the execution of the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_describeExecutionCmd).Standalone()

		transfer_describeExecutionCmd.Flags().String("execution-id", "", "A unique identifier for the execution of a workflow.")
		transfer_describeExecutionCmd.Flags().String("workflow-id", "", "A unique identifier for the workflow.")
		transfer_describeExecutionCmd.MarkFlagRequired("execution-id")
		transfer_describeExecutionCmd.MarkFlagRequired("workflow-id")
	})
	transferCmd.AddCommand(transfer_describeExecutionCmd)
}
