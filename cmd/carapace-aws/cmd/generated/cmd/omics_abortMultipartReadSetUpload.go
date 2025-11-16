package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_abortMultipartReadSetUploadCmd = &cobra.Command{
	Use:   "abort-multipart-read-set-upload",
	Short: "Stops a multipart read set upload into a sequence store and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_abortMultipartReadSetUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_abortMultipartReadSetUploadCmd).Standalone()

		omics_abortMultipartReadSetUploadCmd.Flags().String("sequence-store-id", "", "The sequence store ID for the store involved in the multipart upload.")
		omics_abortMultipartReadSetUploadCmd.Flags().String("upload-id", "", "The ID for the multipart upload.")
		omics_abortMultipartReadSetUploadCmd.MarkFlagRequired("sequence-store-id")
		omics_abortMultipartReadSetUploadCmd.MarkFlagRequired("upload-id")
	})
	omicsCmd.AddCommand(omics_abortMultipartReadSetUploadCmd)
}
