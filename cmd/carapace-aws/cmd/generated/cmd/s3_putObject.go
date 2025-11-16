package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putObjectCmd = &cobra.Command{
	Use:   "put-object",
	Short: "End of support notice: As of October 1, 2025, Amazon S3 has discontinued support for Email Grantee Access Control Lists (ACLs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putObjectCmd).Standalone()

		s3_putObjectCmd.Flags().String("acl", "", "The canned ACL to apply to the object.")
		s3_putObjectCmd.Flags().String("body", "", "Object data.")
		s3_putObjectCmd.Flags().String("bucket", "", "The bucket name to which the PUT action was initiated.")
		s3_putObjectCmd.Flags().String("bucket-key-enabled", "", "Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).")
		s3_putObjectCmd.Flags().String("cache-control", "", "Can be used to specify caching behavior along the request/reply chain.")
		s3_putObjectCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
		s3_putObjectCmd.Flags().String("checksum-crc32", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_putObjectCmd.Flags().String("checksum-crc32-c", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_putObjectCmd.Flags().String("checksum-crc64-nvme", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_putObjectCmd.Flags().String("checksum-sha1", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_putObjectCmd.Flags().String("checksum-sha256", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
		s3_putObjectCmd.Flags().String("content-disposition", "", "Specifies presentational information for the object.")
		s3_putObjectCmd.Flags().String("content-encoding", "", "Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.")
		s3_putObjectCmd.Flags().String("content-language", "", "The language the content is in.")
		s3_putObjectCmd.Flags().String("content-length", "", "Size of the body in bytes.")
		s3_putObjectCmd.Flags().String("content-md5", "", "The Base64 encoded 128-bit `MD5` digest of the message (without the headers) according to RFC 1864. This header can be used as a message integrity check to verify that the data is the same data that was originally sent.")
		s3_putObjectCmd.Flags().String("content-type", "", "A standard MIME type describing the format of the contents.")
		s3_putObjectCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putObjectCmd.Flags().String("expires", "", "The date and time at which the object is no longer cacheable.")
		s3_putObjectCmd.Flags().String("grant-full-control", "", "Gives the grantee READ, READ\\_ACP, and WRITE\\_ACP permissions on the object.")
		s3_putObjectCmd.Flags().String("grant-read", "", "Allows grantee to read the object data and its metadata.")
		s3_putObjectCmd.Flags().String("grant-read-acp", "", "Allows grantee to read the object ACL.")
		s3_putObjectCmd.Flags().String("grant-write-acp", "", "Allows grantee to write the ACL for the applicable object.")
		s3_putObjectCmd.Flags().String("if-match", "", "Uploads the object only if the ETag (entity tag) value provided during the WRITE operation matches the ETag of the object in S3.")
		s3_putObjectCmd.Flags().String("if-none-match", "", "Uploads the object only if the object key name does not already exist in the bucket specified.")
		s3_putObjectCmd.Flags().String("key", "", "Object key for which the PUT action was initiated.")
		s3_putObjectCmd.Flags().String("metadata", "", "A map of metadata to store with the object in S3.")
		s3_putObjectCmd.Flags().String("object-lock-legal-hold-status", "", "Specifies whether a legal hold will be applied to this object.")
		s3_putObjectCmd.Flags().String("object-lock-mode", "", "The Object Lock mode that you want to apply to this object.")
		s3_putObjectCmd.Flags().String("object-lock-retain-until-date", "", "The date and time when you want this object's Object Lock to expire.")
		s3_putObjectCmd.Flags().String("request-payer", "", "")
		s3_putObjectCmd.Flags().String("server-side-encryption", "", "The server-side encryption algorithm that was used when you store this object in Amazon S3 or Amazon FSx.")
		s3_putObjectCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when encrypting the object (for example, `AES256`).")
		s3_putObjectCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.")
		s3_putObjectCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
		s3_putObjectCmd.Flags().String("ssekmsencryption-context", "", "Specifies the Amazon Web Services KMS Encryption Context as an additional encryption context to use for object encryption.")
		s3_putObjectCmd.Flags().String("ssekmskey-id", "", "Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.")
		s3_putObjectCmd.Flags().String("storage-class", "", "By default, Amazon S3 uses the STANDARD Storage Class to store newly created objects.")
		s3_putObjectCmd.Flags().String("tagging", "", "The tag-set for the object.")
		s3_putObjectCmd.Flags().String("website-redirect-location", "", "If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL.")
		s3_putObjectCmd.Flags().String("write-offset-bytes", "", "Specifies the offset for appending data to existing objects in bytes.")
		s3_putObjectCmd.MarkFlagRequired("bucket")
		s3_putObjectCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_putObjectCmd)
}
