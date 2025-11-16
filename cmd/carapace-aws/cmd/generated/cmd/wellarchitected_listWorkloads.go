package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listWorkloadsCmd = &cobra.Command{
	Use:   "list-workloads",
	Short: "Paginated list of workloads.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listWorkloadsCmd).Standalone()

	wellarchitected_listWorkloadsCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
	wellarchitected_listWorkloadsCmd.Flags().String("next-token", "", "")
	wellarchitected_listWorkloadsCmd.Flags().String("workload-name-prefix", "", "")
	wellarchitectedCmd.AddCommand(wellarchitected_listWorkloadsCmd)
}
