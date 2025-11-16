package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createSequenceStoreCmd = &cobra.Command{
	Use:   "create-sequence-store",
	Short: "Creates a sequence store and returns its metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createSequenceStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createSequenceStoreCmd).Standalone()

		omics_createSequenceStoreCmd.Flags().String("client-token", "", "An idempotency token used to dedupe retry requests so that duplicate runs are not created.")
		omics_createSequenceStoreCmd.Flags().String("description", "", "A description for the store.")
		omics_createSequenceStoreCmd.Flags().String("e-tag-algorithm-family", "", "The ETag algorithm family to use for ingested read sets.")
		omics_createSequenceStoreCmd.Flags().String("fallback-location", "", "An S3 location that is used to store files that have failed a direct upload.")
		omics_createSequenceStoreCmd.Flags().String("name", "", "A name for the store.")
		omics_createSequenceStoreCmd.Flags().String("propagated-set-level-tags", "", "The tags keys to propagate to the S3 objects associated with read sets in the sequence store.")
		omics_createSequenceStoreCmd.Flags().String("s3-access-config", "", "S3 access configuration parameters.")
		omics_createSequenceStoreCmd.Flags().String("sse-config", "", "Server-side encryption (SSE) settings for the store.")
		omics_createSequenceStoreCmd.Flags().String("tags", "", "Tags for the store.")
		omics_createSequenceStoreCmd.MarkFlagRequired("name")
	})
	omicsCmd.AddCommand(omics_createSequenceStoreCmd)
}
