package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_deleteVectorEnrichmentJobCmd = &cobra.Command{
	Use:   "delete-vector-enrichment-job",
	Short: "Use this operation to delete a Vector Enrichment job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_deleteVectorEnrichmentJobCmd).Standalone()

	sagemakerGeospatial_deleteVectorEnrichmentJobCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Vector Enrichment job being deleted.")
	sagemakerGeospatial_deleteVectorEnrichmentJobCmd.MarkFlagRequired("arn")
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_deleteVectorEnrichmentJobCmd)
}
