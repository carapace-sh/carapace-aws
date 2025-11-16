package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationFsxLustreCmd = &cobra.Command{
	Use:   "describe-location-fsx-lustre",
	Short: "Provides details about how an DataSync transfer location for an Amazon FSx for Lustre file system is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationFsxLustreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_describeLocationFsxLustreCmd).Standalone()

		datasync_describeLocationFsxLustreCmd.Flags().String("location-arn", "", "The Amazon Resource Name (ARN) of the FSx for Lustre location to describe.")
		datasync_describeLocationFsxLustreCmd.MarkFlagRequired("location-arn")
	})
	datasyncCmd.AddCommand(datasync_describeLocationFsxLustreCmd)
}
