package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_deleteStorageVirtualMachineCmd = &cobra.Command{
	Use:   "delete-storage-virtual-machine",
	Short: "Deletes an existing Amazon FSx for ONTAP storage virtual machine (SVM).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_deleteStorageVirtualMachineCmd).Standalone()

	fsx_deleteStorageVirtualMachineCmd.Flags().String("client-request-token", "", "")
	fsx_deleteStorageVirtualMachineCmd.Flags().String("storage-virtual-machine-id", "", "The ID of the SVM that you want to delete.")
	fsx_deleteStorageVirtualMachineCmd.MarkFlagRequired("storage-virtual-machine-id")
	fsxCmd.AddCommand(fsx_deleteStorageVirtualMachineCmd)
}
