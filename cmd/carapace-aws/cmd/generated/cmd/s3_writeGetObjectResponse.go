package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_writeGetObjectResponseCmd = &cobra.Command{
	Use:   "write-get-object-response",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_writeGetObjectResponseCmd).Standalone()

	s3_writeGetObjectResponseCmd.Flags().String("accept-ranges", "", "Indicates that a range of bytes was specified.")
	s3_writeGetObjectResponseCmd.Flags().String("body", "", "The object data.")
	s3_writeGetObjectResponseCmd.Flags().String("bucket-key-enabled", "", "Indicates whether the object stored in Amazon S3 uses an S3 bucket key for server-side encryption with Amazon Web Services KMS (SSE-KMS).")
	s3_writeGetObjectResponseCmd.Flags().String("cache-control", "", "Specifies caching behavior along the request/reply chain.")
	s3_writeGetObjectResponseCmd.Flags().String("checksum-crc32", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_writeGetObjectResponseCmd.Flags().String("checksum-crc32-c", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_writeGetObjectResponseCmd.Flags().String("checksum-crc64-nvme", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_writeGetObjectResponseCmd.Flags().String("checksum-sha1", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_writeGetObjectResponseCmd.Flags().String("checksum-sha256", "", "This header can be used as a data integrity check to verify that the data received is the same data that was originally sent.")
	s3_writeGetObjectResponseCmd.Flags().String("content-disposition", "", "Specifies presentational information for the object.")
	s3_writeGetObjectResponseCmd.Flags().String("content-encoding", "", "Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.")
	s3_writeGetObjectResponseCmd.Flags().String("content-language", "", "The language the content is in.")
	s3_writeGetObjectResponseCmd.Flags().String("content-length", "", "The size of the content body in bytes.")
	s3_writeGetObjectResponseCmd.Flags().String("content-range", "", "The portion of the object returned in the response.")
	s3_writeGetObjectResponseCmd.Flags().String("content-type", "", "A standard MIME type describing the format of the object data.")
	s3_writeGetObjectResponseCmd.Flags().String("delete-marker", "", "Specifies whether an object stored in Amazon S3 is (`true`) or is not (`false`) a delete marker.")
	s3_writeGetObjectResponseCmd.Flags().String("error-code", "", "A string that uniquely identifies an error condition.")
	s3_writeGetObjectResponseCmd.Flags().String("error-message", "", "Contains a generic description of the error condition.")
	s3_writeGetObjectResponseCmd.Flags().String("etag", "", "An opaque identifier assigned by a web server to a specific version of a resource found at a URL.")
	s3_writeGetObjectResponseCmd.Flags().String("expiration", "", "If the object expiration is configured (see PUT Bucket lifecycle), the response includes this header.")
	s3_writeGetObjectResponseCmd.Flags().String("expires", "", "The date and time at which the object is no longer cacheable.")
	s3_writeGetObjectResponseCmd.Flags().String("last-modified", "", "The date and time that the object was last modified.")
	s3_writeGetObjectResponseCmd.Flags().String("metadata", "", "A map of metadata to store with the object in S3.")
	s3_writeGetObjectResponseCmd.Flags().String("missing-meta", "", "Set to the number of metadata entries not returned in `x-amz-meta` headers.")
	s3_writeGetObjectResponseCmd.Flags().String("object-lock-legal-hold-status", "", "Indicates whether an object stored in Amazon S3 has an active legal hold.")
	s3_writeGetObjectResponseCmd.Flags().String("object-lock-mode", "", "Indicates whether an object stored in Amazon S3 has Object Lock enabled.")
	s3_writeGetObjectResponseCmd.Flags().String("object-lock-retain-until-date", "", "The date and time when Object Lock is configured to expire.")
	s3_writeGetObjectResponseCmd.Flags().String("parts-count", "", "The count of parts this object has.")
	s3_writeGetObjectResponseCmd.Flags().String("replication-status", "", "Indicates if request involves bucket that is either a source or destination in a Replication rule.")
	s3_writeGetObjectResponseCmd.Flags().String("request-charged", "", "")
	s3_writeGetObjectResponseCmd.Flags().String("request-route", "", "Route prefix to the HTTP URL generated.")
	s3_writeGetObjectResponseCmd.Flags().String("request-token", "", "A single use encrypted token that maps `WriteGetObjectResponse` to the end user `GetObject` request.")
	s3_writeGetObjectResponseCmd.Flags().String("restore", "", "Provides information about object restoration operation and expiration time of the restored object copy.")
	s3_writeGetObjectResponseCmd.Flags().String("server-side-encryption", "", "The server-side encryption algorithm used when storing requested object in Amazon S3 or Amazon FSx.")
	s3_writeGetObjectResponseCmd.Flags().String("ssecustomer-algorithm", "", "Encryption algorithm used if server-side encryption with a customer-provided encryption key was specified for object stored in Amazon S3.")
	s3_writeGetObjectResponseCmd.Flags().String("ssecustomer-key-md5", "", "128-bit MD5 digest of customer-provided encryption key used in Amazon S3 to encrypt data stored in S3.")
	s3_writeGetObjectResponseCmd.Flags().String("ssekmskey-id", "", "If present, specifies the ID (Key ID, Key ARN, or Key Alias) of the Amazon Web Services Key Management Service (Amazon Web Services KMS) symmetric encryption customer managed key that was used for stored in Amazon S3 object.")
	s3_writeGetObjectResponseCmd.Flags().String("status-code", "", "The integer status code for an HTTP response of a corresponding `GetObject` request.")
	s3_writeGetObjectResponseCmd.Flags().String("storage-class", "", "Provides storage class information of the object.")
	s3_writeGetObjectResponseCmd.Flags().String("tag-count", "", "The number of tags, if any, on the object.")
	s3_writeGetObjectResponseCmd.Flags().String("version-id", "", "An ID used to reference a specific version of the object.")
	s3_writeGetObjectResponseCmd.MarkFlagRequired("request-route")
	s3_writeGetObjectResponseCmd.MarkFlagRequired("request-token")
	s3Cmd.AddCommand(s3_writeGetObjectResponseCmd)
}
