package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeStorageVirtualMachinesCmd = &cobra.Command{
	Use:   "describe-storage-virtual-machines",
	Short: "Describes one or more Amazon FSx for NetApp ONTAP storage virtual machines (SVMs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeStorageVirtualMachinesCmd).Standalone()

	fsx_describeStorageVirtualMachinesCmd.Flags().String("filters", "", "Enter a filter name:value pair to view a select set of SVMs.")
	fsx_describeStorageVirtualMachinesCmd.Flags().String("max-results", "", "")
	fsx_describeStorageVirtualMachinesCmd.Flags().String("next-token", "", "")
	fsx_describeStorageVirtualMachinesCmd.Flags().String("storage-virtual-machine-ids", "", "Enter the ID of one or more SVMs that you want to view.")
	fsxCmd.AddCommand(fsx_describeStorageVirtualMachinesCmd)
}
