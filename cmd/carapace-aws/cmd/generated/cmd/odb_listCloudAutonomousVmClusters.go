package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listCloudAutonomousVmClustersCmd = &cobra.Command{
	Use:   "list-cloud-autonomous-vm-clusters",
	Short: "Lists all Autonomous VM clusters in a specified Cloud Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listCloudAutonomousVmClustersCmd).Standalone()

	odb_listCloudAutonomousVmClustersCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Cloud Exadata Infrastructure that hosts the Autonomous VM clusters to be listed.")
	odb_listCloudAutonomousVmClustersCmd.Flags().String("max-results", "", "The maximum number of items to return per page.")
	odb_listCloudAutonomousVmClustersCmd.Flags().String("next-token", "", "The pagination token to continue listing from.")
	odbCmd.AddCommand(odb_listCloudAutonomousVmClustersCmd)
}
