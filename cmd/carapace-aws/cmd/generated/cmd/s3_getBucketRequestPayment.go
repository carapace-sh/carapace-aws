package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getBucketRequestPaymentCmd = &cobra.Command{
	Use:   "get-bucket-request-payment",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getBucketRequestPaymentCmd).Standalone()

	s3_getBucketRequestPaymentCmd.Flags().String("bucket", "", "The name of the bucket for which to get the payment request configuration")
	s3_getBucketRequestPaymentCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getBucketRequestPaymentCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_getBucketRequestPaymentCmd)
}
