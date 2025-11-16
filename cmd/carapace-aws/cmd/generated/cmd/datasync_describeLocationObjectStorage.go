package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationObjectStorageCmd = &cobra.Command{
	Use:   "describe-location-object-storage",
	Short: "Provides details about how an DataSync transfer location for an object storage system is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationObjectStorageCmd).Standalone()

	datasync_describeLocationObjectStorageCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the object storage system location.")
	datasync_describeLocationObjectStorageCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_describeLocationObjectStorageCmd)
}
