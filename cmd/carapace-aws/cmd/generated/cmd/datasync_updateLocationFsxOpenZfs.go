package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationFsxOpenZfsCmd = &cobra.Command{
	Use:   "update-location-fsx-open-zfs",
	Short: "Modifies the following configuration parameters of the Amazon FSx for OpenZFS transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationFsxOpenZfsCmd).Standalone()

	datasync_updateLocationFsxOpenZfsCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the FSx for OpenZFS transfer location that you're updating.")
	datasync_updateLocationFsxOpenZfsCmd.Flags().String("protocol", "", "")
	datasync_updateLocationFsxOpenZfsCmd.Flags().String("subdirectory", "", "Specifies a subdirectory in the location's path that must begin with `/fsx`.")
	datasync_updateLocationFsxOpenZfsCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationFsxOpenZfsCmd)
}
