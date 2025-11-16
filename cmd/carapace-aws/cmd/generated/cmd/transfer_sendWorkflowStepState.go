package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_sendWorkflowStepStateCmd = &cobra.Command{
	Use:   "send-workflow-step-state",
	Short: "Sends a callback for asynchronous custom steps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_sendWorkflowStepStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_sendWorkflowStepStateCmd).Standalone()

		transfer_sendWorkflowStepStateCmd.Flags().String("execution-id", "", "A unique identifier for the execution of a workflow.")
		transfer_sendWorkflowStepStateCmd.Flags().String("status", "", "Indicates whether the specified step succeeded or failed.")
		transfer_sendWorkflowStepStateCmd.Flags().String("token", "", "Used to distinguish between multiple callbacks for multiple Lambda steps within the same execution.")
		transfer_sendWorkflowStepStateCmd.Flags().String("workflow-id", "", "A unique identifier for the workflow.")
		transfer_sendWorkflowStepStateCmd.MarkFlagRequired("execution-id")
		transfer_sendWorkflowStepStateCmd.MarkFlagRequired("status")
		transfer_sendWorkflowStepStateCmd.MarkFlagRequired("token")
		transfer_sendWorkflowStepStateCmd.MarkFlagRequired("workflow-id")
	})
	transferCmd.AddCommand(transfer_sendWorkflowStepStateCmd)
}
