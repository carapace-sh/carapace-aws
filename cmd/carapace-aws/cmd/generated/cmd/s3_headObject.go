package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_headObjectCmd = &cobra.Command{
	Use:   "head-object",
	Short: "The `HEAD` operation retrieves metadata from an object without returning the object itself.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_headObjectCmd).Standalone()

	s3_headObjectCmd.Flags().String("bucket", "", "The name of the bucket that contains the object.")
	s3_headObjectCmd.Flags().String("checksum-mode", "", "To retrieve the checksum, this parameter must be enabled.")
	s3_headObjectCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_headObjectCmd.Flags().String("if-match", "", "Return the object only if its entity tag (ETag) is the same as the one specified; otherwise, return a 412 (precondition failed) error.")
	s3_headObjectCmd.Flags().String("if-modified-since", "", "Return the object only if it has been modified since the specified time; otherwise, return a 304 (not modified) error.")
	s3_headObjectCmd.Flags().String("if-none-match", "", "Return the object only if its entity tag (ETag) is different from the one specified; otherwise, return a 304 (not modified) error.")
	s3_headObjectCmd.Flags().String("if-unmodified-since", "", "Return the object only if it has not been modified since the specified time; otherwise, return a 412 (precondition failed) error.")
	s3_headObjectCmd.Flags().String("key", "", "The object key.")
	s3_headObjectCmd.Flags().String("part-number", "", "Part number of the object being read.")
	s3_headObjectCmd.Flags().String("range", "", "HeadObject returns only the metadata for an object.")
	s3_headObjectCmd.Flags().String("request-payer", "", "")
	s3_headObjectCmd.Flags().String("response-cache-control", "", "Sets the `Cache-Control` header of the response.")
	s3_headObjectCmd.Flags().String("response-content-disposition", "", "Sets the `Content-Disposition` header of the response.")
	s3_headObjectCmd.Flags().String("response-content-encoding", "", "Sets the `Content-Encoding` header of the response.")
	s3_headObjectCmd.Flags().String("response-content-language", "", "Sets the `Content-Language` header of the response.")
	s3_headObjectCmd.Flags().String("response-content-type", "", "Sets the `Content-Type` header of the response.")
	s3_headObjectCmd.Flags().String("response-expires", "", "Sets the `Expires` header of the response.")
	s3_headObjectCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when encrypting the object (for example, AES256).")
	s3_headObjectCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.")
	s3_headObjectCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
	s3_headObjectCmd.Flags().String("version-id", "", "Version ID used to reference a specific version of the object.")
	s3_headObjectCmd.MarkFlagRequired("bucket")
	s3_headObjectCmd.MarkFlagRequired("key")
	s3Cmd.AddCommand(s3_headObjectCmd)
}
