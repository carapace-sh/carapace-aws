package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createWorkloadShareCmd = &cobra.Command{
	Use:   "create-workload-share",
	Short: "Create a workload share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createWorkloadShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_createWorkloadShareCmd).Standalone()

		wellarchitected_createWorkloadShareCmd.Flags().String("client-request-token", "", "")
		wellarchitected_createWorkloadShareCmd.Flags().String("permission-type", "", "")
		wellarchitected_createWorkloadShareCmd.Flags().String("shared-with", "", "")
		wellarchitected_createWorkloadShareCmd.Flags().String("workload-id", "", "")
		wellarchitected_createWorkloadShareCmd.MarkFlagRequired("client-request-token")
		wellarchitected_createWorkloadShareCmd.MarkFlagRequired("permission-type")
		wellarchitected_createWorkloadShareCmd.MarkFlagRequired("shared-with")
		wellarchitected_createWorkloadShareCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_createWorkloadShareCmd)
}
