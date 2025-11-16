package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteWorkloadCmd = &cobra.Command{
	Use:   "delete-workload",
	Short: "Delete an existing workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteWorkloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_deleteWorkloadCmd).Standalone()

		wellarchitected_deleteWorkloadCmd.Flags().String("client-request-token", "", "")
		wellarchitected_deleteWorkloadCmd.Flags().String("workload-id", "", "")
		wellarchitected_deleteWorkloadCmd.MarkFlagRequired("client-request-token")
		wellarchitected_deleteWorkloadCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_deleteWorkloadCmd)
}
