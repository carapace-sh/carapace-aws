package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationFsxWindowsCmd = &cobra.Command{
	Use:   "update-location-fsx-windows",
	Short: "Modifies the following configuration parameters of the Amazon FSx for Windows File Server transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationFsxWindowsCmd).Standalone()

	datasync_updateLocationFsxWindowsCmd.Flags().String("domain", "", "Specifies the name of the Windows domain that your FSx for Windows File Server file system belongs to.")
	datasync_updateLocationFsxWindowsCmd.Flags().String("location-arn", "", "Specifies the ARN of the FSx for Windows File Server transfer location that you're updating.")
	datasync_updateLocationFsxWindowsCmd.Flags().String("password", "", "Specifies the password of the user with the permissions to mount and access the files, folders, and file metadata in your FSx for Windows File Server file system.")
	datasync_updateLocationFsxWindowsCmd.Flags().String("subdirectory", "", "Specifies a mount path for your file system using forward slashes.")
	datasync_updateLocationFsxWindowsCmd.Flags().String("user", "", "Specifies the user with the permissions to mount and access the files, folders, and file metadata in your FSx for Windows File Server file system.")
	datasync_updateLocationFsxWindowsCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationFsxWindowsCmd)
}
