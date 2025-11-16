package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteWorkloadShareCmd = &cobra.Command{
	Use:   "delete-workload-share",
	Short: "Delete a workload share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteWorkloadShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_deleteWorkloadShareCmd).Standalone()

		wellarchitected_deleteWorkloadShareCmd.Flags().String("client-request-token", "", "")
		wellarchitected_deleteWorkloadShareCmd.Flags().String("share-id", "", "")
		wellarchitected_deleteWorkloadShareCmd.Flags().String("workload-id", "", "")
		wellarchitected_deleteWorkloadShareCmd.MarkFlagRequired("client-request-token")
		wellarchitected_deleteWorkloadShareCmd.MarkFlagRequired("share-id")
		wellarchitected_deleteWorkloadShareCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_deleteWorkloadShareCmd)
}
