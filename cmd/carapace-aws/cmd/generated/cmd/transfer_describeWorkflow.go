package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeWorkflowCmd = &cobra.Command{
	Use:   "describe-workflow",
	Short: "Describes the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeWorkflowCmd).Standalone()

	transfer_describeWorkflowCmd.Flags().String("workflow-id", "", "A unique identifier for the workflow.")
	transfer_describeWorkflowCmd.MarkFlagRequired("workflow-id")
	transferCmd.AddCommand(transfer_describeWorkflowCmd)
}
