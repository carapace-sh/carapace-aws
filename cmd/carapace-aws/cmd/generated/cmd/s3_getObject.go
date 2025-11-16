package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectCmd = &cobra.Command{
	Use:   "get-object",
	Short: "Retrieves an object from Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getObjectCmd).Standalone()

		s3_getObjectCmd.Flags().String("bucket", "", "The bucket name containing the object.")
		s3_getObjectCmd.Flags().String("checksum-mode", "", "To retrieve the checksum, this mode must be enabled.")
		s3_getObjectCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getObjectCmd.Flags().String("if-match", "", "Return the object only if its entity tag (ETag) is the same as the one specified in this header; otherwise, return a `412 Precondition Failed` error.")
		s3_getObjectCmd.Flags().String("if-modified-since", "", "Return the object only if it has been modified since the specified time; otherwise, return a `304 Not Modified` error.")
		s3_getObjectCmd.Flags().String("if-none-match", "", "Return the object only if its entity tag (ETag) is different from the one specified in this header; otherwise, return a `304 Not Modified` error.")
		s3_getObjectCmd.Flags().String("if-unmodified-since", "", "Return the object only if it has not been modified since the specified time; otherwise, return a `412 Precondition Failed` error.")
		s3_getObjectCmd.Flags().String("key", "", "Key of the object to get.")
		s3_getObjectCmd.Flags().String("part-number", "", "Part number of the object being read.")
		s3_getObjectCmd.Flags().String("range", "", "Downloads the specified byte range of an object.")
		s3_getObjectCmd.Flags().String("request-payer", "", "")
		s3_getObjectCmd.Flags().String("response-cache-control", "", "Sets the `Cache-Control` header of the response.")
		s3_getObjectCmd.Flags().String("response-content-disposition", "", "Sets the `Content-Disposition` header of the response.")
		s3_getObjectCmd.Flags().String("response-content-encoding", "", "Sets the `Content-Encoding` header of the response.")
		s3_getObjectCmd.Flags().String("response-content-language", "", "Sets the `Content-Language` header of the response.")
		s3_getObjectCmd.Flags().String("response-content-type", "", "Sets the `Content-Type` header of the response.")
		s3_getObjectCmd.Flags().String("response-expires", "", "Sets the `Expires` header of the response.")
		s3_getObjectCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when decrypting the object (for example, `AES256`).")
		s3_getObjectCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key that you originally provided for Amazon S3 to encrypt the data before storing it.")
		s3_getObjectCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the customer-provided encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
		s3_getObjectCmd.Flags().String("version-id", "", "Version ID used to reference a specific version of the object.")
		s3_getObjectCmd.MarkFlagRequired("bucket")
		s3_getObjectCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_getObjectCmd)
}
