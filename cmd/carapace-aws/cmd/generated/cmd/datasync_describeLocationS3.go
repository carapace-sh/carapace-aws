package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationS3Cmd = &cobra.Command{
	Use:   "describe-location-s3",
	Short: "Provides details about how an DataSync transfer location for an S3 bucket is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationS3Cmd).Standalone()

	datasync_describeLocationS3Cmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the Amazon S3 location.")
	datasync_describeLocationS3Cmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_describeLocationS3Cmd)
}
