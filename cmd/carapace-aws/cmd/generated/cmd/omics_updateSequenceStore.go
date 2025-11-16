package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateSequenceStoreCmd = &cobra.Command{
	Use:   "update-sequence-store",
	Short: "Update one or more parameters for the sequence store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateSequenceStoreCmd).Standalone()

	omics_updateSequenceStoreCmd.Flags().String("client-token", "", "To ensure that requests don't run multiple times, specify a unique token for each request.")
	omics_updateSequenceStoreCmd.Flags().String("description", "", "A description for the sequence store.")
	omics_updateSequenceStoreCmd.Flags().String("fallback-location", "", "The S3 URI of a bucket and folder to store Read Sets that fail to upload.")
	omics_updateSequenceStoreCmd.Flags().String("id", "", "The ID of the sequence store.")
	omics_updateSequenceStoreCmd.Flags().String("name", "", "A name for the sequence store.")
	omics_updateSequenceStoreCmd.Flags().String("propagated-set-level-tags", "", "The tags keys to propagate to the S3 objects associated with read sets in the sequence store.")
	omics_updateSequenceStoreCmd.Flags().String("s3-access-config", "", "S3 access configuration parameters.")
	omics_updateSequenceStoreCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_updateSequenceStoreCmd)
}
