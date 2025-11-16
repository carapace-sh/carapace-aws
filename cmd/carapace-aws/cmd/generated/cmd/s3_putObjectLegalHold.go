package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putObjectLegalHoldCmd = &cobra.Command{
	Use:   "put-object-legal-hold",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putObjectLegalHoldCmd).Standalone()

	s3_putObjectLegalHoldCmd.Flags().String("bucket", "", "The bucket name containing the object that you want to place a legal hold on.")
	s3_putObjectLegalHoldCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
	s3_putObjectLegalHoldCmd.Flags().String("content-md5", "", "The MD5 hash for the request body.")
	s3_putObjectLegalHoldCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putObjectLegalHoldCmd.Flags().String("key", "", "The key name for the object that you want to place a legal hold on.")
	s3_putObjectLegalHoldCmd.Flags().String("legal-hold", "", "Container element for the legal hold configuration you want to apply to the specified object.")
	s3_putObjectLegalHoldCmd.Flags().String("request-payer", "", "")
	s3_putObjectLegalHoldCmd.Flags().String("version-id", "", "The version ID of the object that you want to place a legal hold on.")
	s3_putObjectLegalHoldCmd.MarkFlagRequired("bucket")
	s3_putObjectLegalHoldCmd.MarkFlagRequired("key")
	s3Cmd.AddCommand(s3_putObjectLegalHoldCmd)
}
