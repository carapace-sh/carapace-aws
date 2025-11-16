package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationFsxWindowsCmd = &cobra.Command{
	Use:   "create-location-fsx-windows",
	Short: "Creates a transfer *location* for an Amazon FSx for Windows File Server file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationFsxWindowsCmd).Standalone()

	datasync_createLocationFsxWindowsCmd.Flags().String("domain", "", "Specifies the name of the Windows domain that the FSx for Windows File Server file system belongs to.")
	datasync_createLocationFsxWindowsCmd.Flags().String("fsx-filesystem-arn", "", "Specifies the Amazon Resource Name (ARN) for the FSx for Windows File Server file system.")
	datasync_createLocationFsxWindowsCmd.Flags().String("password", "", "Specifies the password of the user with the permissions to mount and access the files, folders, and file metadata in your FSx for Windows File Server file system.")
	datasync_createLocationFsxWindowsCmd.Flags().String("security-group-arns", "", "Specifies the ARNs of the Amazon EC2 security groups that provide access to your file system's preferred subnet.")
	datasync_createLocationFsxWindowsCmd.Flags().String("subdirectory", "", "Specifies a mount path for your file system using forward slashes.")
	datasync_createLocationFsxWindowsCmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
	datasync_createLocationFsxWindowsCmd.Flags().String("user", "", "Specifies the user with the permissions to mount and access the files, folders, and file metadata in your FSx for Windows File Server file system.")
	datasync_createLocationFsxWindowsCmd.MarkFlagRequired("fsx-filesystem-arn")
	datasync_createLocationFsxWindowsCmd.MarkFlagRequired("password")
	datasync_createLocationFsxWindowsCmd.MarkFlagRequired("security-group-arns")
	datasync_createLocationFsxWindowsCmd.MarkFlagRequired("user")
	datasyncCmd.AddCommand(datasync_createLocationFsxWindowsCmd)
}
