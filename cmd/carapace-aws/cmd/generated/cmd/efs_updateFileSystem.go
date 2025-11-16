package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_updateFileSystemCmd = &cobra.Command{
	Use:   "update-file-system",
	Short: "Updates the throughput mode or the amount of provisioned throughput of an existing file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_updateFileSystemCmd).Standalone()

	efs_updateFileSystemCmd.Flags().String("file-system-id", "", "The ID of the file system that you want to update.")
	efs_updateFileSystemCmd.Flags().String("provisioned-throughput-in-mibps", "", "(Optional) The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating.")
	efs_updateFileSystemCmd.Flags().String("throughput-mode", "", "(Optional) Updates the file system's throughput mode.")
	efs_updateFileSystemCmd.MarkFlagRequired("file-system-id")
	efsCmd.AddCommand(efs_updateFileSystemCmd)
}
