package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationFsxOpenZfsCmd = &cobra.Command{
	Use:   "create-location-fsx-open-zfs",
	Short: "Creates a transfer *location* for an Amazon FSx for OpenZFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationFsxOpenZfsCmd).Standalone()

	datasync_createLocationFsxOpenZfsCmd.Flags().String("fsx-filesystem-arn", "", "The Amazon Resource Name (ARN) of the FSx for OpenZFS file system.")
	datasync_createLocationFsxOpenZfsCmd.Flags().String("protocol", "", "The type of protocol that DataSync uses to access your file system.")
	datasync_createLocationFsxOpenZfsCmd.Flags().String("security-group-arns", "", "The ARNs of the security groups that are used to configure the FSx for OpenZFS file system.")
	datasync_createLocationFsxOpenZfsCmd.Flags().String("subdirectory", "", "A subdirectory in the location's path that must begin with `/fsx`.")
	datasync_createLocationFsxOpenZfsCmd.Flags().String("tags", "", "The key-value pair that represents a tag that you want to add to the resource.")
	datasync_createLocationFsxOpenZfsCmd.MarkFlagRequired("fsx-filesystem-arn")
	datasync_createLocationFsxOpenZfsCmd.MarkFlagRequired("protocol")
	datasync_createLocationFsxOpenZfsCmd.MarkFlagRequired("security-group-arns")
	datasyncCmd.AddCommand(datasync_createLocationFsxOpenZfsCmd)
}
