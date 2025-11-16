package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listPartsCmd = &cobra.Command{
	Use:   "list-parts",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listPartsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_listPartsCmd).Standalone()

		s3_listPartsCmd.Flags().String("bucket", "", "The name of the bucket to which the parts are being uploaded.")
		s3_listPartsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_listPartsCmd.Flags().String("key", "", "Object key for which the multipart upload was initiated.")
		s3_listPartsCmd.Flags().String("max-parts", "", "Sets the maximum number of parts to return.")
		s3_listPartsCmd.Flags().String("part-number-marker", "", "Specifies the part after which listing should begin.")
		s3_listPartsCmd.Flags().String("request-payer", "", "")
		s3_listPartsCmd.Flags().String("ssecustomer-algorithm", "", "The server-side encryption (SSE) algorithm used to encrypt the object.")
		s3_listPartsCmd.Flags().String("ssecustomer-key", "", "The server-side encryption (SSE) customer managed key.")
		s3_listPartsCmd.Flags().String("ssecustomer-key-md5", "", "The MD5 server-side encryption (SSE) customer managed key.")
		s3_listPartsCmd.Flags().String("upload-id", "", "Upload ID identifying the multipart upload whose parts are being listed.")
		s3_listPartsCmd.MarkFlagRequired("bucket")
		s3_listPartsCmd.MarkFlagRequired("key")
		s3_listPartsCmd.MarkFlagRequired("upload-id")
	})
	s3Cmd.AddCommand(s3_listPartsCmd)
}
