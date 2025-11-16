package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_createMultipartUploadCmd = &cobra.Command{
	Use:   "create-multipart-upload",
	Short: "End of support notice: As of October 1, 2025, Amazon S3 has discontinued support for Email Grantee Access Control Lists (ACLs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_createMultipartUploadCmd).Standalone()

	s3_createMultipartUploadCmd.Flags().String("acl", "", "The canned ACL to apply to the object.")
	s3_createMultipartUploadCmd.Flags().String("bucket", "", "The name of the bucket where the multipart upload is initiated and where the object is uploaded.")
	s3_createMultipartUploadCmd.Flags().String("bucket-key-enabled", "", "Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).")
	s3_createMultipartUploadCmd.Flags().String("cache-control", "", "Specifies caching behavior along the request/reply chain.")
	s3_createMultipartUploadCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm that you want Amazon S3 to use to create the checksum for the object.")
	s3_createMultipartUploadCmd.Flags().String("checksum-type", "", "Indicates the checksum type that you want Amazon S3 to use to calculate the object’s checksum value.")
	s3_createMultipartUploadCmd.Flags().String("content-disposition", "", "Specifies presentational information for the object.")
	s3_createMultipartUploadCmd.Flags().String("content-encoding", "", "Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.")
	s3_createMultipartUploadCmd.Flags().String("content-language", "", "The language that the content is in.")
	s3_createMultipartUploadCmd.Flags().String("content-type", "", "A standard MIME type describing the format of the object data.")
	s3_createMultipartUploadCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_createMultipartUploadCmd.Flags().String("expires", "", "The date and time at which the object is no longer cacheable.")
	s3_createMultipartUploadCmd.Flags().String("grant-full-control", "", "Specify access permissions explicitly to give the grantee READ, READ\\_ACP, and WRITE\\_ACP permissions on the object.")
	s3_createMultipartUploadCmd.Flags().String("grant-read", "", "Specify access permissions explicitly to allow grantee to read the object data and its metadata.")
	s3_createMultipartUploadCmd.Flags().String("grant-read-acp", "", "Specify access permissions explicitly to allows grantee to read the object ACL.")
	s3_createMultipartUploadCmd.Flags().String("grant-write-acp", "", "Specify access permissions explicitly to allows grantee to allow grantee to write the ACL for the applicable object.")
	s3_createMultipartUploadCmd.Flags().String("key", "", "Object key for which the multipart upload is to be initiated.")
	s3_createMultipartUploadCmd.Flags().String("metadata", "", "A map of metadata to store with the object in S3.")
	s3_createMultipartUploadCmd.Flags().String("object-lock-legal-hold-status", "", "Specifies whether you want to apply a legal hold to the uploaded object.")
	s3_createMultipartUploadCmd.Flags().String("object-lock-mode", "", "Specifies the Object Lock mode that you want to apply to the uploaded object.")
	s3_createMultipartUploadCmd.Flags().String("object-lock-retain-until-date", "", "Specifies the date and time when you want the Object Lock to expire.")
	s3_createMultipartUploadCmd.Flags().String("request-payer", "", "")
	s3_createMultipartUploadCmd.Flags().String("server-side-encryption", "", "The server-side encryption algorithm used when you store this object in Amazon S3 or Amazon FSx.")
	s3_createMultipartUploadCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when encrypting the object (for example, AES256).")
	s3_createMultipartUploadCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.")
	s3_createMultipartUploadCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the customer-provided encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
	s3_createMultipartUploadCmd.Flags().String("ssekmsencryption-context", "", "Specifies the Amazon Web Services KMS Encryption Context to use for object encryption.")
	s3_createMultipartUploadCmd.Flags().String("ssekmskey-id", "", "Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.")
	s3_createMultipartUploadCmd.Flags().String("storage-class", "", "By default, Amazon S3 uses the STANDARD Storage Class to store newly created objects.")
	s3_createMultipartUploadCmd.Flags().String("tagging", "", "The tag-set for the object.")
	s3_createMultipartUploadCmd.Flags().String("website-redirect-location", "", "If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL.")
	s3_createMultipartUploadCmd.MarkFlagRequired("bucket")
	s3_createMultipartUploadCmd.MarkFlagRequired("key")
	s3Cmd.AddCommand(s3_createMultipartUploadCmd)
}
