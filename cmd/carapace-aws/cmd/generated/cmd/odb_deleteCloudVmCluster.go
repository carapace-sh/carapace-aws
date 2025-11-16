package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_deleteCloudVmClusterCmd = &cobra.Command{
	Use:   "delete-cloud-vm-cluster",
	Short: "Deletes the specified VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_deleteCloudVmClusterCmd).Standalone()

	odb_deleteCloudVmClusterCmd.Flags().String("cloud-vm-cluster-id", "", "The unique identifier of the VM cluster to delete.")
	odb_deleteCloudVmClusterCmd.MarkFlagRequired("cloud-vm-cluster-id")
	odbCmd.AddCommand(odb_deleteCloudVmClusterCmd)
}
