package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReadSetMetadataCmd = &cobra.Command{
	Use:   "get-read-set-metadata",
	Short: "Retrieves the metadata for a read set from a sequence store in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReadSetMetadataCmd).Standalone()

	omics_getReadSetMetadataCmd.Flags().String("id", "", "The read set's ID.")
	omics_getReadSetMetadataCmd.Flags().String("sequence-store-id", "", "The read set's sequence store ID.")
	omics_getReadSetMetadataCmd.MarkFlagRequired("id")
	omics_getReadSetMetadataCmd.MarkFlagRequired("sequence-store-id")
	omicsCmd.AddCommand(omics_getReadSetMetadataCmd)
}
