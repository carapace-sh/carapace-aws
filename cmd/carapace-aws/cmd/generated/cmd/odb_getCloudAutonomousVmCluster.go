package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getCloudAutonomousVmClusterCmd = &cobra.Command{
	Use:   "get-cloud-autonomous-vm-cluster",
	Short: "Gets information about a specific Autonomous VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getCloudAutonomousVmClusterCmd).Standalone()

	odb_getCloudAutonomousVmClusterCmd.Flags().String("cloud-autonomous-vm-cluster-id", "", "The unique identifier of the Autonomous VM cluster to retrieve information about.")
	odb_getCloudAutonomousVmClusterCmd.MarkFlagRequired("cloud-autonomous-vm-cluster-id")
	odbCmd.AddCommand(odb_getCloudAutonomousVmClusterCmd)
}
