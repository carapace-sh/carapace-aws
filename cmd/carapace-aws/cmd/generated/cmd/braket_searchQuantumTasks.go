package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_searchQuantumTasksCmd = &cobra.Command{
	Use:   "search-quantum-tasks",
	Short: "Searches for tasks that match the specified filter values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_searchQuantumTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_searchQuantumTasksCmd).Standalone()

		braket_searchQuantumTasksCmd.Flags().String("filters", "", "Array of `SearchQuantumTasksFilter` objects to use when searching for quantum tasks.")
		braket_searchQuantumTasksCmd.Flags().String("max-results", "", "Maximum number of results to return in the response.")
		braket_searchQuantumTasksCmd.Flags().String("next-token", "", "A token used for pagination of results returned in the response.")
		braket_searchQuantumTasksCmd.MarkFlagRequired("filters")
	})
	braketCmd.AddCommand(braket_searchQuantumTasksCmd)
}
