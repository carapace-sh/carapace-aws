package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReadSetUploadPartsCmd = &cobra.Command{
	Use:   "list-read-set-upload-parts",
	Short: "Lists all parts in a multipart read set upload for a sequence store and returns the metadata in a JSON formatted output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReadSetUploadPartsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listReadSetUploadPartsCmd).Standalone()

		omics_listReadSetUploadPartsCmd.Flags().String("filter", "", "Attributes used to filter for a specific subset of read set part uploads.")
		omics_listReadSetUploadPartsCmd.Flags().String("max-results", "", "The maximum number of read set upload parts returned in a page.")
		omics_listReadSetUploadPartsCmd.Flags().String("next-token", "", "Next token returned in the response of a previous ListReadSetUploadPartsRequest call.")
		omics_listReadSetUploadPartsCmd.Flags().String("part-source", "", "The source file for the upload part.")
		omics_listReadSetUploadPartsCmd.Flags().String("sequence-store-id", "", "The Sequence Store ID used for the multipart uploads.")
		omics_listReadSetUploadPartsCmd.Flags().String("upload-id", "", "The ID for the initiated multipart upload.")
		omics_listReadSetUploadPartsCmd.MarkFlagRequired("part-source")
		omics_listReadSetUploadPartsCmd.MarkFlagRequired("sequence-store-id")
		omics_listReadSetUploadPartsCmd.MarkFlagRequired("upload-id")
	})
	omicsCmd.AddCommand(omics_listReadSetUploadPartsCmd)
}
