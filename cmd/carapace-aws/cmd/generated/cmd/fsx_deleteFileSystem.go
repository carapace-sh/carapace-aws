package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_deleteFileSystemCmd = &cobra.Command{
	Use:   "delete-file-system",
	Short: "Deletes a file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_deleteFileSystemCmd).Standalone()

	fsx_deleteFileSystemCmd.Flags().String("client-request-token", "", "A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent deletion.")
	fsx_deleteFileSystemCmd.Flags().String("file-system-id", "", "The ID of the file system that you want to delete.")
	fsx_deleteFileSystemCmd.Flags().String("lustre-configuration", "", "")
	fsx_deleteFileSystemCmd.Flags().String("open-zfsconfiguration", "", "The configuration object for the OpenZFS file system used in the `DeleteFileSystem` operation.")
	fsx_deleteFileSystemCmd.Flags().String("windows-configuration", "", "")
	fsx_deleteFileSystemCmd.MarkFlagRequired("file-system-id")
	fsxCmd.AddCommand(fsx_deleteFileSystemCmd)
}
