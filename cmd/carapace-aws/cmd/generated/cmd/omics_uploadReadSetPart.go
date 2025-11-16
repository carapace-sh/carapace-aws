package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_uploadReadSetPartCmd = &cobra.Command{
	Use:   "upload-read-set-part",
	Short: "Uploads a specific part of a read set into a sequence store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_uploadReadSetPartCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_uploadReadSetPartCmd).Standalone()

		omics_uploadReadSetPartCmd.Flags().String("part-number", "", "The number of the part being uploaded.")
		omics_uploadReadSetPartCmd.Flags().String("part-source", "", "The source file for an upload part.")
		omics_uploadReadSetPartCmd.Flags().String("payload", "", "The read set data to upload for a part.")
		omics_uploadReadSetPartCmd.Flags().String("sequence-store-id", "", "The Sequence Store ID used for the multipart upload.")
		omics_uploadReadSetPartCmd.Flags().String("upload-id", "", "The ID for the initiated multipart upload.")
		omics_uploadReadSetPartCmd.MarkFlagRequired("part-number")
		omics_uploadReadSetPartCmd.MarkFlagRequired("part-source")
		omics_uploadReadSetPartCmd.MarkFlagRequired("payload")
		omics_uploadReadSetPartCmd.MarkFlagRequired("sequence-store-id")
		omics_uploadReadSetPartCmd.MarkFlagRequired("upload-id")
	})
	omicsCmd.AddCommand(omics_uploadReadSetPartCmd)
}
