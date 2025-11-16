package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_completeMultipartReadSetUploadCmd = &cobra.Command{
	Use:   "complete-multipart-read-set-upload",
	Short: "Completes a multipart read set upload into a sequence store after you have initiated the upload process with `CreateMultipartReadSetUpload` and uploaded all read set parts using `UploadReadSetPart`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_completeMultipartReadSetUploadCmd).Standalone()

	omics_completeMultipartReadSetUploadCmd.Flags().String("parts", "", "The individual uploads or parts of a multipart upload.")
	omics_completeMultipartReadSetUploadCmd.Flags().String("sequence-store-id", "", "The sequence store ID for the store involved in the multipart upload.")
	omics_completeMultipartReadSetUploadCmd.Flags().String("upload-id", "", "The ID for the multipart upload.")
	omics_completeMultipartReadSetUploadCmd.MarkFlagRequired("parts")
	omics_completeMultipartReadSetUploadCmd.MarkFlagRequired("sequence-store-id")
	omics_completeMultipartReadSetUploadCmd.MarkFlagRequired("upload-id")
	omicsCmd.AddCommand(omics_completeMultipartReadSetUploadCmd)
}
