package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listWorkloadSharesCmd = &cobra.Command{
	Use:   "list-workload-shares",
	Short: "List the workload shares associated with the workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listWorkloadSharesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_listWorkloadSharesCmd).Standalone()

		wellarchitected_listWorkloadSharesCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
		wellarchitected_listWorkloadSharesCmd.Flags().String("next-token", "", "")
		wellarchitected_listWorkloadSharesCmd.Flags().String("shared-with-prefix", "", "The Amazon Web Services account ID, organization ID, or organizational unit (OU) ID with which the workload is shared.")
		wellarchitected_listWorkloadSharesCmd.Flags().String("status", "", "")
		wellarchitected_listWorkloadSharesCmd.Flags().String("workload-id", "", "")
		wellarchitected_listWorkloadSharesCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_listWorkloadSharesCmd)
}
