package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_getQuantumTaskCmd = &cobra.Command{
	Use:   "get-quantum-task",
	Short: "Retrieves the specified quantum task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_getQuantumTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_getQuantumTaskCmd).Standalone()

		braket_getQuantumTaskCmd.Flags().String("additional-attribute-names", "", "A list of attributes to return additional information for.")
		braket_getQuantumTaskCmd.Flags().String("quantum-task-arn", "", "The ARN of the quantum task to retrieve.")
		braket_getQuantumTaskCmd.MarkFlagRequired("quantum-task-arn")
	})
	braketCmd.AddCommand(braket_getQuantumTaskCmd)
}
