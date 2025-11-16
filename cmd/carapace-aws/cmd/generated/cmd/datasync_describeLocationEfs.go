package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationEfsCmd = &cobra.Command{
	Use:   "describe-location-efs",
	Short: "Provides details about how an DataSync transfer location for an Amazon EFS file system is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationEfsCmd).Standalone()

	datasync_describeLocationEfsCmd.Flags().String("location-arn", "", "The Amazon Resource Name (ARN) of the Amazon EFS file system location that you want information about.")
	datasync_describeLocationEfsCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_describeLocationEfsCmd)
}
