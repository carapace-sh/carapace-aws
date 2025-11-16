package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getReferenceMetadataCmd = &cobra.Command{
	Use:   "get-reference-metadata",
	Short: "Retrieves metadata for a reference genome.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getReferenceMetadataCmd).Standalone()

	omics_getReferenceMetadataCmd.Flags().String("id", "", "The reference's ID.")
	omics_getReferenceMetadataCmd.Flags().String("reference-store-id", "", "The reference's reference store ID.")
	omics_getReferenceMetadataCmd.MarkFlagRequired("id")
	omics_getReferenceMetadataCmd.MarkFlagRequired("reference-store-id")
	omicsCmd.AddCommand(omics_getReferenceMetadataCmd)
}
