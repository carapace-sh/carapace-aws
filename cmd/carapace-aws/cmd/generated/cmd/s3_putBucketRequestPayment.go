package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putBucketRequestPaymentCmd = &cobra.Command{
	Use:   "put-bucket-request-payment",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putBucketRequestPaymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putBucketRequestPaymentCmd).Standalone()

		s3_putBucketRequestPaymentCmd.Flags().String("bucket", "", "The bucket name.")
		s3_putBucketRequestPaymentCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the request when you use the SDK.")
		s3_putBucketRequestPaymentCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the data.")
		s3_putBucketRequestPaymentCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putBucketRequestPaymentCmd.Flags().String("request-payment-configuration", "", "Container for Payer.")
		s3_putBucketRequestPaymentCmd.MarkFlagRequired("bucket")
		s3_putBucketRequestPaymentCmd.MarkFlagRequired("request-payment-configuration")
	})
	s3Cmd.AddCommand(s3_putBucketRequestPaymentCmd)
}
