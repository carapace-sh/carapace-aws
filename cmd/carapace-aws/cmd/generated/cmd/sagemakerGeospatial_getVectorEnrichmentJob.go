package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_getVectorEnrichmentJobCmd = &cobra.Command{
	Use:   "get-vector-enrichment-job",
	Short: "Retrieves details of a Vector Enrichment Job for a given job Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_getVectorEnrichmentJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_getVectorEnrichmentJobCmd).Standalone()

		sagemakerGeospatial_getVectorEnrichmentJobCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Vector Enrichment job.")
		sagemakerGeospatial_getVectorEnrichmentJobCmd.MarkFlagRequired("arn")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_getVectorEnrichmentJobCmd)
}
