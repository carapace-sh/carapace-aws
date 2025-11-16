package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createStorageVirtualMachineCmd = &cobra.Command{
	Use:   "create-storage-virtual-machine",
	Short: "Creates a storage virtual machine (SVM) for an Amazon FSx for ONTAP file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createStorageVirtualMachineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_createStorageVirtualMachineCmd).Standalone()

		fsx_createStorageVirtualMachineCmd.Flags().String("active-directory-configuration", "", "Describes the self-managed Microsoft Active Directory to which you want to join the SVM.")
		fsx_createStorageVirtualMachineCmd.Flags().String("client-request-token", "", "")
		fsx_createStorageVirtualMachineCmd.Flags().String("file-system-id", "", "")
		fsx_createStorageVirtualMachineCmd.Flags().String("name", "", "The name of the SVM.")
		fsx_createStorageVirtualMachineCmd.Flags().String("root-volume-security-style", "", "The security style of the root volume of the SVM.")
		fsx_createStorageVirtualMachineCmd.Flags().String("svm-admin-password", "", "The password to use when managing the SVM using the NetApp ONTAP CLI or REST API.")
		fsx_createStorageVirtualMachineCmd.Flags().String("tags", "", "")
		fsx_createStorageVirtualMachineCmd.MarkFlagRequired("file-system-id")
		fsx_createStorageVirtualMachineCmd.MarkFlagRequired("name")
	})
	fsxCmd.AddCommand(fsx_createStorageVirtualMachineCmd)
}
