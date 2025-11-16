package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_deleteVirtualClusterCmd = &cobra.Command{
	Use:   "delete-virtual-cluster",
	Short: "Deletes a virtual cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_deleteVirtualClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_deleteVirtualClusterCmd).Standalone()

		emrContainers_deleteVirtualClusterCmd.Flags().String("id", "", "The ID of the virtual cluster that will be deleted.")
		emrContainers_deleteVirtualClusterCmd.MarkFlagRequired("id")
	})
	emrContainersCmd.AddCommand(emrContainers_deleteVirtualClusterCmd)
}
