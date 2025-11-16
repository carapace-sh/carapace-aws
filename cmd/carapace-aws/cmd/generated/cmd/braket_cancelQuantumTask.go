package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_cancelQuantumTaskCmd = &cobra.Command{
	Use:   "cancel-quantum-task",
	Short: "Cancels the specified task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_cancelQuantumTaskCmd).Standalone()

	braket_cancelQuantumTaskCmd.Flags().String("client-token", "", "The client token associated with the cancellation request.")
	braket_cancelQuantumTaskCmd.Flags().String("quantum-task-arn", "", "The ARN of the quantum task to cancel.")
	braket_cancelQuantumTaskCmd.MarkFlagRequired("client-token")
	braket_cancelQuantumTaskCmd.MarkFlagRequired("quantum-task-arn")
	braketCmd.AddCommand(braket_cancelQuantumTaskCmd)
}
