package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_completeMultipartUploadCmd = &cobra.Command{
	Use:   "complete-multipart-upload",
	Short: "Completes a multipart upload by assembling previously uploaded parts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_completeMultipartUploadCmd).Standalone()

	s3_completeMultipartUploadCmd.Flags().String("bucket", "", "Name of the bucket to which the multipart upload was initiated.")
	s3_completeMultipartUploadCmd.Flags().String("checksum-crc32", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_completeMultipartUploadCmd.Flags().String("checksum-crc32-c", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_completeMultipartUploadCmd.Flags().String("checksum-crc64-nvme", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_completeMultipartUploadCmd.Flags().String("checksum-sha1", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_completeMultipartUploadCmd.Flags().String("checksum-sha256", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_completeMultipartUploadCmd.Flags().String("checksum-type", "", "This header specifies the checksum type of the object, which determines how part-level checksums are combined to create an object-level checksum for multipart objects.")
	s3_completeMultipartUploadCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_completeMultipartUploadCmd.Flags().String("if-match", "", "Uploads the object only if the ETag (entity tag) value provided during the WRITE operation matches the ETag of the object in S3.")
	s3_completeMultipartUploadCmd.Flags().String("if-none-match", "", "Uploads the object only if the object key name does not already exist in the bucket specified.")
	s3_completeMultipartUploadCmd.Flags().String("key", "", "Object key for which the multipart upload was initiated.")
	s3_completeMultipartUploadCmd.Flags().String("mpu-object-size", "", "The expected total object size of the multipart upload request.")
	s3_completeMultipartUploadCmd.Flags().String("multipart-upload", "", "The container for the multipart upload request information.")
	s3_completeMultipartUploadCmd.Flags().String("request-payer", "", "")
	s3_completeMultipartUploadCmd.Flags().String("ssecustomer-algorithm", "", "The server-side encryption (SSE) algorithm used to encrypt the object.")
	s3_completeMultipartUploadCmd.Flags().String("ssecustomer-key", "", "The server-side encryption (SSE) customer managed key.")
	s3_completeMultipartUploadCmd.Flags().String("ssecustomer-key-md5", "", "The MD5 server-side encryption (SSE) customer managed key.")
	s3_completeMultipartUploadCmd.Flags().String("upload-id", "", "ID for the initiated multipart upload.")
	s3_completeMultipartUploadCmd.MarkFlagRequired("bucket")
	s3_completeMultipartUploadCmd.MarkFlagRequired("key")
	s3_completeMultipartUploadCmd.MarkFlagRequired("upload-id")
	s3Cmd.AddCommand(s3_completeMultipartUploadCmd)
}
