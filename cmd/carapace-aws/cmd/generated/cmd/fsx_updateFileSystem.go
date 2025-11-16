package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_updateFileSystemCmd = &cobra.Command{
	Use:   "update-file-system",
	Short: "Use this operation to update the configuration of an existing Amazon FSx file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_updateFileSystemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_updateFileSystemCmd).Standalone()

		fsx_updateFileSystemCmd.Flags().String("client-request-token", "", "A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent updates.")
		fsx_updateFileSystemCmd.Flags().String("file-system-id", "", "The ID of the file system that you are updating.")
		fsx_updateFileSystemCmd.Flags().String("file-system-type-version", "", "The Lustre version you are updating an FSx for Lustre file system to.")
		fsx_updateFileSystemCmd.Flags().String("lustre-configuration", "", "")
		fsx_updateFileSystemCmd.Flags().String("network-type", "", "Changes the network type of an FSx for OpenZFS file system.")
		fsx_updateFileSystemCmd.Flags().String("ontap-configuration", "", "")
		fsx_updateFileSystemCmd.Flags().String("open-zfsconfiguration", "", "The configuration updates for an FSx for OpenZFS file system.")
		fsx_updateFileSystemCmd.Flags().String("storage-capacity", "", "Use this parameter to increase the storage capacity of an FSx for Windows File Server, FSx for Lustre, FSx for OpenZFS, or FSx for ONTAP file system.")
		fsx_updateFileSystemCmd.Flags().String("storage-type", "", "")
		fsx_updateFileSystemCmd.Flags().String("windows-configuration", "", "The configuration updates for an Amazon FSx for Windows File Server file system.")
		fsx_updateFileSystemCmd.MarkFlagRequired("file-system-id")
	})
	fsxCmd.AddCommand(fsx_updateFileSystemCmd)
}
