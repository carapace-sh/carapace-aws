package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_deleteCloudAutonomousVmClusterCmd = &cobra.Command{
	Use:   "delete-cloud-autonomous-vm-cluster",
	Short: "Deletes an Autonomous VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_deleteCloudAutonomousVmClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_deleteCloudAutonomousVmClusterCmd).Standalone()

		odb_deleteCloudAutonomousVmClusterCmd.Flags().String("cloud-autonomous-vm-cluster-id", "", "The unique identifier of the Autonomous VM cluster to delete.")
		odb_deleteCloudAutonomousVmClusterCmd.MarkFlagRequired("cloud-autonomous-vm-cluster-id")
	})
	odbCmd.AddCommand(odb_deleteCloudAutonomousVmClusterCmd)
}
