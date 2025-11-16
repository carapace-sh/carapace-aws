package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putPublicAccessBlockCmd = &cobra.Command{
	Use:   "put-public-access-block",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putPublicAccessBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putPublicAccessBlockCmd).Standalone()

		s3_putPublicAccessBlockCmd.Flags().String("bucket", "", "The name of the Amazon S3 bucket whose `PublicAccessBlock` configuration you want to set.")
		s3_putPublicAccessBlockCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
		s3_putPublicAccessBlockCmd.Flags().String("content-md5", "", "The MD5 hash of the `PutPublicAccessBlock` request body.")
		s3_putPublicAccessBlockCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putPublicAccessBlockCmd.Flags().String("public-access-block-configuration", "", "The `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket.")
		s3_putPublicAccessBlockCmd.MarkFlagRequired("bucket")
		s3_putPublicAccessBlockCmd.MarkFlagRequired("public-access-block-configuration")
	})
	s3Cmd.AddCommand(s3_putPublicAccessBlockCmd)
}
