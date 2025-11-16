package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createFileSystemCmd = &cobra.Command{
	Use:   "create-file-system",
	Short: "Creates a new, empty Amazon FSx file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createFileSystemCmd).Standalone()

	fsx_createFileSystemCmd.Flags().String("client-request-token", "", "A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent creation.")
	fsx_createFileSystemCmd.Flags().String("file-system-type", "", "The type of Amazon FSx file system to create.")
	fsx_createFileSystemCmd.Flags().String("file-system-type-version", "", "For FSx for Lustre file systems, sets the Lustre version for the file system that you're creating.")
	fsx_createFileSystemCmd.Flags().String("kms-key-id", "", "")
	fsx_createFileSystemCmd.Flags().String("lustre-configuration", "", "")
	fsx_createFileSystemCmd.Flags().String("network-type", "", "The network type of the Amazon FSx file system that you are creating.")
	fsx_createFileSystemCmd.Flags().String("ontap-configuration", "", "")
	fsx_createFileSystemCmd.Flags().String("open-zfsconfiguration", "", "The OpenZFS configuration for the file system that's being created.")
	fsx_createFileSystemCmd.Flags().String("security-group-ids", "", "A list of IDs specifying the security groups to apply to all network interfaces created for file system access.")
	fsx_createFileSystemCmd.Flags().String("storage-capacity", "", "Sets the storage capacity of the file system that you're creating, in gibibytes (GiB).")
	fsx_createFileSystemCmd.Flags().String("storage-type", "", "Sets the storage class for the file system that you're creating.")
	fsx_createFileSystemCmd.Flags().String("subnet-ids", "", "Specifies the IDs of the subnets that the file system will be accessible from.")
	fsx_createFileSystemCmd.Flags().String("tags", "", "The tags to apply to the file system that's being created.")
	fsx_createFileSystemCmd.Flags().String("windows-configuration", "", "The Microsoft Windows configuration for the file system that's being created.")
	fsx_createFileSystemCmd.MarkFlagRequired("file-system-type")
	fsx_createFileSystemCmd.MarkFlagRequired("subnet-ids")
	fsxCmd.AddCommand(fsx_createFileSystemCmd)
}
