package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationFsxOpenZfsCmd = &cobra.Command{
	Use:   "describe-location-fsx-open-zfs",
	Short: "Provides details about how an DataSync transfer location for an Amazon FSx for OpenZFS file system is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationFsxOpenZfsCmd).Standalone()

	datasync_describeLocationFsxOpenZfsCmd.Flags().String("location-arn", "", "The Amazon Resource Name (ARN) of the FSx for OpenZFS location to describe.")
	datasync_describeLocationFsxOpenZfsCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_describeLocationFsxOpenZfsCmd)
}
