package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listMultipartUploadsCmd = &cobra.Command{
	Use:   "list-multipart-uploads",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listMultipartUploadsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_listMultipartUploadsCmd).Standalone()

		s3_listMultipartUploadsCmd.Flags().String("bucket", "", "The name of the bucket to which the multipart upload was initiated.")
		s3_listMultipartUploadsCmd.Flags().String("delimiter", "", "Character you use to group keys.")
		s3_listMultipartUploadsCmd.Flags().String("encoding-type", "", "")
		s3_listMultipartUploadsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_listMultipartUploadsCmd.Flags().String("key-marker", "", "Specifies the multipart upload after which listing should begin.")
		s3_listMultipartUploadsCmd.Flags().String("max-uploads", "", "Sets the maximum number of multipart uploads, from 1 to 1,000, to return in the response body.")
		s3_listMultipartUploadsCmd.Flags().String("prefix", "", "Lists in-progress uploads only for those keys that begin with the specified prefix.")
		s3_listMultipartUploadsCmd.Flags().String("request-payer", "", "")
		s3_listMultipartUploadsCmd.Flags().String("upload-id-marker", "", "Together with key-marker, specifies the multipart upload after which listing should begin.")
		s3_listMultipartUploadsCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_listMultipartUploadsCmd)
}
