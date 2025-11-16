package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getWorkloadCmd = &cobra.Command{
	Use:   "get-workload",
	Short: "Get an existing workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getWorkloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_getWorkloadCmd).Standalone()

		wellarchitected_getWorkloadCmd.Flags().String("workload-id", "", "")
		wellarchitected_getWorkloadCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_getWorkloadCmd)
}
