package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectAttributesCmd = &cobra.Command{
	Use:   "get-object-attributes",
	Short: "Retrieves all of the metadata from an object without returning the object itself.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getObjectAttributesCmd).Standalone()

		s3_getObjectAttributesCmd.Flags().String("bucket", "", "The name of the bucket that contains the object.")
		s3_getObjectAttributesCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getObjectAttributesCmd.Flags().String("key", "", "The object key.")
		s3_getObjectAttributesCmd.Flags().String("max-parts", "", "Sets the maximum number of parts to return.")
		s3_getObjectAttributesCmd.Flags().String("object-attributes", "", "Specifies the fields at the root level that you want returned in the response.")
		s3_getObjectAttributesCmd.Flags().String("part-number-marker", "", "Specifies the part after which listing should begin.")
		s3_getObjectAttributesCmd.Flags().String("request-payer", "", "")
		s3_getObjectAttributesCmd.Flags().String("ssecustomer-algorithm", "", "Specifies the algorithm to use when encrypting the object (for example, AES256).")
		s3_getObjectAttributesCmd.Flags().String("ssecustomer-key", "", "Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data.")
		s3_getObjectAttributesCmd.Flags().String("ssecustomer-key-md5", "", "Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.")
		s3_getObjectAttributesCmd.Flags().String("version-id", "", "The version ID used to reference a specific version of the object.")
		s3_getObjectAttributesCmd.MarkFlagRequired("bucket")
		s3_getObjectAttributesCmd.MarkFlagRequired("key")
		s3_getObjectAttributesCmd.MarkFlagRequired("object-attributes")
	})
	s3Cmd.AddCommand(s3_getObjectAttributesCmd)
}
