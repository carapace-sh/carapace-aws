package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_uploadPartCopyCmd = &cobra.Command{
	Use:   "upload-part-copy",
	Short: "Uploads a part by copying data from an existing object as data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_uploadPartCopyCmd).Standalone()

	s3_uploadPartCopyCmd.Flags().String("bucket", "", "The bucket name.")
	s3_uploadPartCopyCmd.Flags().String("copy-source", "", "Specifies the source object for the copy operation.")
	s3_uploadPartCopyCmd.Flags().String("copy-source-if-match", "", "Copies the object if its entity tag (ETag) matches the specified tag.")
	s3_uploadPartCopyCmd.Flags().String("copy-source-if-modified-since", "", "Copies the object if it has been modified since the specified time.")
	s3_uploadPartCopyCmd.Flags().String("copy-source-if-none-match", "", "Copies the object if its entity tag (ETag) is different than the specified ETag.")
	s3_uploadPartCopyCmd.Flags().String("copy-source-if-unmodified-since", "", "Copies the object if it hasn't been modified since the specified time.")
	s3_uploadPartCopyCmd.Flags().String("copy-source-range", "", "The range of bytes to copy from the source object.")
	s3_uploadPartCopyCmd.Flags().String("copy-source-ssecustomer-algorithm", "", "Specifies the algorithm to use when decrypting the source object (for example, `AES256`).")
	s3_uploadPartCopyCmd.Flags().String("copy-source-ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object.")
	s3_uploadPartCopyCmd.Flags().String("copy-source-ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
	s3_uploadPartCopyCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected destination bucket owner.")
	s3_uploadPartCopyCmd.Flags().String("expected-source-bucket-owner", "", "The account ID of the expected source bucket owner.")
	s3_uploadPartCopyCmd.Flags().String("key", "", "Object key for which the multipart upload was initiated.")
	s3_uploadPartCopyCmd.Flags().String("part-number", "", "Part number of part being copied.")
	s3_uploadPartCopyCmd.Flags().String("request-payer", "", "")
	s3_uploadPartCopyCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when encrypting the object (for example, AES256).")
	s3_uploadPartCopyCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.")
	s3_uploadPartCopyCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
	s3_uploadPartCopyCmd.Flags().String("upload-id", "", "Upload ID identifying the multipart upload whose part is being copied.")
	s3_uploadPartCopyCmd.MarkFlagRequired("bucket")
	s3_uploadPartCopyCmd.MarkFlagRequired("copy-source")
	s3_uploadPartCopyCmd.MarkFlagRequired("key")
	s3_uploadPartCopyCmd.MarkFlagRequired("part-number")
	s3_uploadPartCopyCmd.MarkFlagRequired("upload-id")
	s3Cmd.AddCommand(s3_uploadPartCopyCmd)
}
