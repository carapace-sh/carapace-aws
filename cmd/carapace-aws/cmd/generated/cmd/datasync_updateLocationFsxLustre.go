package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationFsxLustreCmd = &cobra.Command{
	Use:   "update-location-fsx-lustre",
	Short: "Modifies the following configuration parameters of the Amazon FSx for Lustre transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationFsxLustreCmd).Standalone()

	datasync_updateLocationFsxLustreCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the FSx for Lustre transfer location that you're updating.")
	datasync_updateLocationFsxLustreCmd.Flags().String("subdirectory", "", "Specifies a mount path for your FSx for Lustre file system.")
	datasync_updateLocationFsxLustreCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationFsxLustreCmd)
}
