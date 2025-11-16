package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_uploadPartCmd = &cobra.Command{
	Use:   "upload-part",
	Short: "Uploads a part in a multipart upload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_uploadPartCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_uploadPartCmd).Standalone()

		s3_uploadPartCmd.Flags().String("body", "", "Object data.")
		s3_uploadPartCmd.Flags().String("bucket", "", "The name of the bucket to which the multipart upload was initiated.")
		s3_uploadPartCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
		s3_uploadPartCmd.Flags().String("checksum-crc32", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_uploadPartCmd.Flags().String("checksum-crc32-c", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_uploadPartCmd.Flags().String("checksum-crc64-nvme", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_uploadPartCmd.Flags().String("checksum-sha1", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_uploadPartCmd.Flags().String("checksum-sha256", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_uploadPartCmd.Flags().String("content-length", "", "Size of the body in bytes.")
		s3_uploadPartCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit MD5 digest of the part data.")
		s3_uploadPartCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_uploadPartCmd.Flags().String("key", "", "Object key for which the multipart upload was initiated.")
		s3_uploadPartCmd.Flags().String("part-number", "", "Part number of part being uploaded.")
		s3_uploadPartCmd.Flags().String("request-payer", "", "")
		s3_uploadPartCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when encrypting the object (for example, AES256).")
		s3_uploadPartCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.")
		s3_uploadPartCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
		s3_uploadPartCmd.Flags().String("upload-id", "", "Upload ID identifying the multipart upload whose part is being uploaded.")
		s3_uploadPartCmd.MarkFlagRequired("bucket")
		s3_uploadPartCmd.MarkFlagRequired("key")
		s3_uploadPartCmd.MarkFlagRequired("part-number")
		s3_uploadPartCmd.MarkFlagRequired("upload-id")
	})
	s3Cmd.AddCommand(s3_uploadPartCmd)
}
