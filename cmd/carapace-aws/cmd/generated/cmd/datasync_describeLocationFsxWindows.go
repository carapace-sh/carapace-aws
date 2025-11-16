package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationFsxWindowsCmd = &cobra.Command{
	Use:   "describe-location-fsx-windows",
	Short: "Provides details about how an DataSync transfer location for an Amazon FSx for Windows File Server file system is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationFsxWindowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_describeLocationFsxWindowsCmd).Standalone()

		datasync_describeLocationFsxWindowsCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the FSx for Windows File Server location.")
		datasync_describeLocationFsxWindowsCmd.MarkFlagRequired("location-arn")
	})
	datasyncCmd.AddCommand(datasync_describeLocationFsxWindowsCmd)
}
