package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationS3Cmd = &cobra.Command{
	Use:   "update-location-s3",
	Short: "Modifies the following configuration parameters of the Amazon S3 transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationS3Cmd).Standalone()

	datasync_updateLocationS3Cmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the Amazon S3 transfer location that you're updating.")
	datasync_updateLocationS3Cmd.Flags().String("s3-config", "", "")
	datasync_updateLocationS3Cmd.Flags().String("s3-storage-class", "", "Specifies the storage class that you want your objects to use when Amazon S3 is a transfer destination.")
	datasync_updateLocationS3Cmd.Flags().String("subdirectory", "", "Specifies a prefix in the S3 bucket that DataSync reads from or writes to (depending on whether the bucket is a source or destination location).")
	datasync_updateLocationS3Cmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationS3Cmd)
}
