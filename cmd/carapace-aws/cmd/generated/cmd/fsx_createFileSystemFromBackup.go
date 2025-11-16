package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createFileSystemFromBackupCmd = &cobra.Command{
	Use:   "create-file-system-from-backup",
	Short: "Creates a new Amazon FSx for Lustre, Amazon FSx for Windows File Server, or Amazon FSx for OpenZFS file system from an existing Amazon FSx backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createFileSystemFromBackupCmd).Standalone()

	fsx_createFileSystemFromBackupCmd.Flags().String("backup-id", "", "")
	fsx_createFileSystemFromBackupCmd.Flags().String("client-request-token", "", "A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent creation.")
	fsx_createFileSystemFromBackupCmd.Flags().String("file-system-type-version", "", "Sets the version for the Amazon FSx for Lustre file system that you're creating from a backup.")
	fsx_createFileSystemFromBackupCmd.Flags().String("kms-key-id", "", "")
	fsx_createFileSystemFromBackupCmd.Flags().String("lustre-configuration", "", "")
	fsx_createFileSystemFromBackupCmd.Flags().String("network-type", "", "Sets the network type for the Amazon FSx for OpenZFS file system that you're creating from a backup.")
	fsx_createFileSystemFromBackupCmd.Flags().String("open-zfsconfiguration", "", "The OpenZFS configuration for the file system that's being created.")
	fsx_createFileSystemFromBackupCmd.Flags().String("security-group-ids", "", "A list of IDs for the security groups that apply to the specified network interfaces created for file system access.")
	fsx_createFileSystemFromBackupCmd.Flags().String("storage-capacity", "", "Sets the storage capacity of the OpenZFS file system that you're creating from a backup, in gibibytes (GiB).")
	fsx_createFileSystemFromBackupCmd.Flags().String("storage-type", "", "Sets the storage type for the Windows, OpenZFS, or Lustre file system that you're creating from a backup.")
	fsx_createFileSystemFromBackupCmd.Flags().String("subnet-ids", "", "Specifies the IDs of the subnets that the file system will be accessible from.")
	fsx_createFileSystemFromBackupCmd.Flags().String("tags", "", "The tags to be applied to the file system at file system creation.")
	fsx_createFileSystemFromBackupCmd.Flags().String("windows-configuration", "", "The configuration for this Microsoft Windows file system.")
	fsx_createFileSystemFromBackupCmd.MarkFlagRequired("backup-id")
	fsx_createFileSystemFromBackupCmd.MarkFlagRequired("subnet-ids")
	fsxCmd.AddCommand(fsx_createFileSystemFromBackupCmd)
}
