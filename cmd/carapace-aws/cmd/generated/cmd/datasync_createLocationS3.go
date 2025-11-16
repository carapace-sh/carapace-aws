package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationS3Cmd = &cobra.Command{
	Use:   "create-location-s3",
	Short: "Creates a transfer *location* for an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationS3Cmd).Standalone()

	datasync_createLocationS3Cmd.Flags().String("agent-arns", "", "(Amazon S3 on Outposts only) Specifies the Amazon Resource Name (ARN) of the DataSync agent on your Outpost.")
	datasync_createLocationS3Cmd.Flags().String("s3-bucket-arn", "", "Specifies the ARN of the S3 bucket that you want to use as a location.")
	datasync_createLocationS3Cmd.Flags().String("s3-config", "", "")
	datasync_createLocationS3Cmd.Flags().String("s3-storage-class", "", "Specifies the storage class that you want your objects to use when Amazon S3 is a transfer destination.")
	datasync_createLocationS3Cmd.Flags().String("subdirectory", "", "Specifies a prefix in the S3 bucket that DataSync reads from or writes to (depending on whether the bucket is a source or destination location).")
	datasync_createLocationS3Cmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
	datasync_createLocationS3Cmd.MarkFlagRequired("s3-bucket-arn")
	datasync_createLocationS3Cmd.MarkFlagRequired("s3-config")
	datasyncCmd.AddCommand(datasync_createLocationS3Cmd)
}
