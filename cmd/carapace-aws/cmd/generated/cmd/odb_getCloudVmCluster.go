package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getCloudVmClusterCmd = &cobra.Command{
	Use:   "get-cloud-vm-cluster",
	Short: "Returns information about the specified VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getCloudVmClusterCmd).Standalone()

	odb_getCloudVmClusterCmd.Flags().String("cloud-vm-cluster-id", "", "The unique identifier of the VM cluster.")
	odb_getCloudVmClusterCmd.MarkFlagRequired("cloud-vm-cluster-id")
	odbCmd.AddCommand(odb_getCloudVmClusterCmd)
}
