package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_releaseFileSystemNfsV3LocksCmd = &cobra.Command{
	Use:   "release-file-system-nfs-v3-locks",
	Short: "Releases the file system lock from an Amazon FSx for OpenZFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_releaseFileSystemNfsV3LocksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_releaseFileSystemNfsV3LocksCmd).Standalone()

		fsx_releaseFileSystemNfsV3LocksCmd.Flags().String("client-request-token", "", "")
		fsx_releaseFileSystemNfsV3LocksCmd.Flags().String("file-system-id", "", "")
		fsx_releaseFileSystemNfsV3LocksCmd.MarkFlagRequired("file-system-id")
	})
	fsxCmd.AddCommand(fsx_releaseFileSystemNfsV3LocksCmd)
}
