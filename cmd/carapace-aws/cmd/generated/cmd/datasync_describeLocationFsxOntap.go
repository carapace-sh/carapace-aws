package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationFsxOntapCmd = &cobra.Command{
	Use:   "describe-location-fsx-ontap",
	Short: "Provides details about how an DataSync transfer location for an Amazon FSx for NetApp ONTAP file system is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationFsxOntapCmd).Standalone()

	datasync_describeLocationFsxOntapCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the FSx for ONTAP file system location that you want information about.")
	datasync_describeLocationFsxOntapCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_describeLocationFsxOntapCmd)
}
