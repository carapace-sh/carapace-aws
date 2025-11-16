package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_updateStorageVirtualMachineCmd = &cobra.Command{
	Use:   "update-storage-virtual-machine",
	Short: "Updates an FSx for ONTAP storage virtual machine (SVM).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_updateStorageVirtualMachineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_updateStorageVirtualMachineCmd).Standalone()

		fsx_updateStorageVirtualMachineCmd.Flags().String("active-directory-configuration", "", "Specifies updates to an SVM's Microsoft Active Directory (AD) configuration.")
		fsx_updateStorageVirtualMachineCmd.Flags().String("client-request-token", "", "")
		fsx_updateStorageVirtualMachineCmd.Flags().String("storage-virtual-machine-id", "", "The ID of the SVM that you want to update, in the format `svm-0123456789abcdef0`.")
		fsx_updateStorageVirtualMachineCmd.Flags().String("svm-admin-password", "", "Specifies a new SvmAdminPassword.")
		fsx_updateStorageVirtualMachineCmd.MarkFlagRequired("storage-virtual-machine-id")
	})
	fsxCmd.AddCommand(fsx_updateStorageVirtualMachineCmd)
}
