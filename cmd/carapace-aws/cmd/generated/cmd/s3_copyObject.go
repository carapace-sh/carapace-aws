package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_copyObjectCmd = &cobra.Command{
	Use:   "copy-object",
	Short: "Creates a copy of an object that is already stored in Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_copyObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_copyObjectCmd).Standalone()

		s3_copyObjectCmd.Flags().String("acl", "", "The canned access control list (ACL) to apply to the object.")
		s3_copyObjectCmd.Flags().String("bucket", "", "The name of the destination bucket.")
		s3_copyObjectCmd.Flags().String("bucket-key-enabled", "", "Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).")
		s3_copyObjectCmd.Flags().String("cache-control", "", "Specifies the caching behavior along the request/reply chain.")
		s3_copyObjectCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm that you want Amazon S3 to use to create the checksum for the object.")
		s3_copyObjectCmd.Flags().String("content-disposition", "", "Specifies presentational information for the object.")
		s3_copyObjectCmd.Flags().String("content-encoding", "", "Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.")
		s3_copyObjectCmd.Flags().String("content-language", "", "The language the content is in.")
		s3_copyObjectCmd.Flags().String("content-type", "", "A standard MIME type that describes the format of the object data.")
		s3_copyObjectCmd.Flags().String("copy-source", "", "Specifies the source object for the copy operation.")
		s3_copyObjectCmd.Flags().String("copy-source-if-match", "", "Copies the object if its entity tag (ETag) matches the specified tag.")
		s3_copyObjectCmd.Flags().String("copy-source-if-modified-since", "", "Copies the object if it has been modified since the specified time.")
		s3_copyObjectCmd.Flags().String("copy-source-if-none-match", "", "Copies the object if its entity tag (ETag) is different than the specified ETag.")
		s3_copyObjectCmd.Flags().String("copy-source-if-unmodified-since", "", "Copies the object if it hasn't been modified since the specified time.")
		s3_copyObjectCmd.Flags().String("copy-source-ssecustomer-algorithm", "", "Specifies the algorithm to use when decrypting the source object (for example, `AES256`).")
		s3_copyObjectCmd.Flags().String("copy-source-ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object.")
		s3_copyObjectCmd.Flags().String("copy-source-ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
		s3_copyObjectCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected destination bucket owner.")
		s3_copyObjectCmd.Flags().String("expected-source-bucket-owner", "", "The account ID of the expected source bucket owner.")
		s3_copyObjectCmd.Flags().String("expires", "", "The date and time at which the object is no longer cacheable.")
		s3_copyObjectCmd.Flags().String("grant-full-control", "", "Gives the grantee READ, READ\\_ACP, and WRITE\\_ACP permissions on the object.")
		s3_copyObjectCmd.Flags().String("grant-read", "", "Allows grantee to read the object data and its metadata.")
		s3_copyObjectCmd.Flags().String("grant-read-acp", "", "Allows grantee to read the object ACL.")
		s3_copyObjectCmd.Flags().String("grant-write-acp", "", "Allows grantee to write the ACL for the applicable object.")
		s3_copyObjectCmd.Flags().String("if-match", "", "Copies the object if the entity tag (ETag) of the destination object matches the specified tag.")
		s3_copyObjectCmd.Flags().String("if-none-match", "", "Copies the object only if the object key name at the destination does not already exist in the bucket specified.")
		s3_copyObjectCmd.Flags().String("key", "", "The key of the destination object.")
		s3_copyObjectCmd.Flags().String("metadata", "", "A map of metadata to store with the object in S3.")
		s3_copyObjectCmd.Flags().String("metadata-directive", "", "Specifies whether the metadata is copied from the source object or replaced with metadata that's provided in the request.")
		s3_copyObjectCmd.Flags().String("object-lock-legal-hold-status", "", "Specifies whether you want to apply a legal hold to the object copy.")
		s3_copyObjectCmd.Flags().String("object-lock-mode", "", "The Object Lock mode that you want to apply to the object copy.")
		s3_copyObjectCmd.Flags().String("object-lock-retain-until-date", "", "The date and time when you want the Object Lock of the object copy to expire.")
		s3_copyObjectCmd.Flags().String("request-payer", "", "")
		s3_copyObjectCmd.Flags().String("server-side-encryption", "", "The server-side encryption algorithm used when storing this object in Amazon S3.")
		s3_copyObjectCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when encrypting the object (for example, `AES256`).")
		s3_copyObjectCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.")
		s3_copyObjectCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
		s3_copyObjectCmd.Flags().String("ssekmsencryption-context", "", "Specifies the Amazon Web Services KMS Encryption Context as an additional encryption context to use for the destination object encryption.")
		s3_copyObjectCmd.Flags().String("ssekmskey-id", "", "Specifies the KMS key ID (Key ID, Key ARN, or Key Alias) to use for object encryption.")
		s3_copyObjectCmd.Flags().String("storage-class", "", "If the `x-amz-storage-class` header is not used, the copied object will be stored in the `STANDARD` Storage Class by default.")
		s3_copyObjectCmd.Flags().String("tagging", "", "The tag-set for the object copy in the destination bucket.")
		s3_copyObjectCmd.Flags().String("tagging-directive", "", "Specifies whether the object tag-set is copied from the source object or replaced with the tag-set that's provided in the request.")
		s3_copyObjectCmd.Flags().String("website-redirect-location", "", "If the destination bucket is configured as a website, redirects requests for this object copy to another object in the same bucket or to an external URL.")
		s3_copyObjectCmd.MarkFlagRequired("bucket")
		s3_copyObjectCmd.MarkFlagRequired("copy-source")
		s3_copyObjectCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_copyObjectCmd)
}
