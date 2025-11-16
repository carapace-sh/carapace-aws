package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_stopVectorEnrichmentJobCmd = &cobra.Command{
	Use:   "stop-vector-enrichment-job",
	Short: "Stops the Vector Enrichment job for a given job ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_stopVectorEnrichmentJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_stopVectorEnrichmentJobCmd).Standalone()

		sagemakerGeospatial_stopVectorEnrichmentJobCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Vector Enrichment job.")
		sagemakerGeospatial_stopVectorEnrichmentJobCmd.MarkFlagRequired("arn")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_stopVectorEnrichmentJobCmd)
}
