package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationFsxLustreCmd = &cobra.Command{
	Use:   "create-location-fsx-lustre",
	Short: "Creates a transfer *location* for an Amazon FSx for Lustre file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationFsxLustreCmd).Standalone()

	datasync_createLocationFsxLustreCmd.Flags().String("fsx-filesystem-arn", "", "Specifies the Amazon Resource Name (ARN) of the FSx for Lustre file system.")
	datasync_createLocationFsxLustreCmd.Flags().String("security-group-arns", "", "Specifies the Amazon Resource Names (ARNs) of up to five security groups that provide access to your FSx for Lustre file system.")
	datasync_createLocationFsxLustreCmd.Flags().String("subdirectory", "", "Specifies a mount path for your FSx for Lustre file system.")
	datasync_createLocationFsxLustreCmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
	datasync_createLocationFsxLustreCmd.MarkFlagRequired("fsx-filesystem-arn")
	datasync_createLocationFsxLustreCmd.MarkFlagRequired("security-group-arns")
	datasyncCmd.AddCommand(datasync_createLocationFsxLustreCmd)
}
