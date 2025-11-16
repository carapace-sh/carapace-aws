package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listMultipartReadSetUploadsCmd = &cobra.Command{
	Use:   "list-multipart-read-set-uploads",
	Short: "Lists in-progress multipart read set uploads for a sequence store and returns it in a JSON formatted output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listMultipartReadSetUploadsCmd).Standalone()

	omics_listMultipartReadSetUploadsCmd.Flags().String("max-results", "", "The maximum number of multipart uploads returned in a page.")
	omics_listMultipartReadSetUploadsCmd.Flags().String("next-token", "", "Next token returned in the response of a previous ListMultipartReadSetUploads call.")
	omics_listMultipartReadSetUploadsCmd.Flags().String("sequence-store-id", "", "The Sequence Store ID used for the multipart uploads.")
	omics_listMultipartReadSetUploadsCmd.MarkFlagRequired("sequence-store-id")
	omicsCmd.AddCommand(omics_listMultipartReadSetUploadsCmd)
}
