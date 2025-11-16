package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateWorkloadShareCmd = &cobra.Command{
	Use:   "update-workload-share",
	Short: "Update a workload share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateWorkloadShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_updateWorkloadShareCmd).Standalone()

		wellarchitected_updateWorkloadShareCmd.Flags().String("permission-type", "", "")
		wellarchitected_updateWorkloadShareCmd.Flags().String("share-id", "", "")
		wellarchitected_updateWorkloadShareCmd.Flags().String("workload-id", "", "")
		wellarchitected_updateWorkloadShareCmd.MarkFlagRequired("permission-type")
		wellarchitected_updateWorkloadShareCmd.MarkFlagRequired("share-id")
		wellarchitected_updateWorkloadShareCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_updateWorkloadShareCmd)
}
