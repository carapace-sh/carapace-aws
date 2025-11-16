package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_abortMultipartUploadCmd = &cobra.Command{
	Use:   "abort-multipart-upload",
	Short: "This operation aborts a multipart upload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_abortMultipartUploadCmd).Standalone()

	s3_abortMultipartUploadCmd.Flags().String("bucket", "", "The bucket name to which the upload was taking place.")
	s3_abortMultipartUploadCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_abortMultipartUploadCmd.Flags().String("if-match-initiated-time", "", "If present, this header aborts an in progress multipart upload only if it was initiated on the provided timestamp.")
	s3_abortMultipartUploadCmd.Flags().String("key", "", "Key of the object for which the multipart upload was initiated.")
	s3_abortMultipartUploadCmd.Flags().String("request-payer", "", "")
	s3_abortMultipartUploadCmd.Flags().String("upload-id", "", "Upload ID that identifies the multipart upload.")
	s3_abortMultipartUploadCmd.MarkFlagRequired("bucket")
	s3_abortMultipartUploadCmd.MarkFlagRequired("key")
	s3_abortMultipartUploadCmd.MarkFlagRequired("upload-id")
	s3Cmd.AddCommand(s3_abortMultipartUploadCmd)
}
