package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createMultipartReadSetUploadCmd = &cobra.Command{
	Use:   "create-multipart-read-set-upload",
	Short: "Initiates a multipart read set upload for uploading partitioned source files into a sequence store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createMultipartReadSetUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createMultipartReadSetUploadCmd).Standalone()

		omics_createMultipartReadSetUploadCmd.Flags().String("client-token", "", "An idempotency token that can be used to avoid triggering multiple multipart uploads.")
		omics_createMultipartReadSetUploadCmd.Flags().String("description", "", "The description of the read set.")
		omics_createMultipartReadSetUploadCmd.Flags().String("generated-from", "", "Where the source originated.")
		omics_createMultipartReadSetUploadCmd.Flags().String("name", "", "The name of the read set.")
		omics_createMultipartReadSetUploadCmd.Flags().String("reference-arn", "", "The ARN of the reference.")
		omics_createMultipartReadSetUploadCmd.Flags().String("sample-id", "", "The source's sample ID.")
		omics_createMultipartReadSetUploadCmd.Flags().String("sequence-store-id", "", "The sequence store ID for the store that is the destination of the multipart uploads.")
		omics_createMultipartReadSetUploadCmd.Flags().String("source-file-type", "", "The type of file being uploaded.")
		omics_createMultipartReadSetUploadCmd.Flags().String("subject-id", "", "The source's subject ID.")
		omics_createMultipartReadSetUploadCmd.Flags().String("tags", "", "Any tags to add to the read set.")
		omics_createMultipartReadSetUploadCmd.MarkFlagRequired("name")
		omics_createMultipartReadSetUploadCmd.MarkFlagRequired("sample-id")
		omics_createMultipartReadSetUploadCmd.MarkFlagRequired("sequence-store-id")
		omics_createMultipartReadSetUploadCmd.MarkFlagRequired("source-file-type")
		omics_createMultipartReadSetUploadCmd.MarkFlagRequired("subject-id")
	})
	omicsCmd.AddCommand(omics_createMultipartReadSetUploadCmd)
}
