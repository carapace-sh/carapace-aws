package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_exportVectorEnrichmentJobCmd = &cobra.Command{
	Use:   "export-vector-enrichment-job",
	Short: "Use this operation to copy results of a Vector Enrichment job to an Amazon S3 location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_exportVectorEnrichmentJobCmd).Standalone()

	sagemakerGeospatial_exportVectorEnrichmentJobCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Vector Enrichment job.")
	sagemakerGeospatial_exportVectorEnrichmentJobCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
	sagemakerGeospatial_exportVectorEnrichmentJobCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM rolewith permission to upload to the location in OutputConfig.")
	sagemakerGeospatial_exportVectorEnrichmentJobCmd.Flags().String("output-config", "", "Output location information for exporting Vector Enrichment Job results.")
	sagemakerGeospatial_exportVectorEnrichmentJobCmd.MarkFlagRequired("arn")
	sagemakerGeospatial_exportVectorEnrichmentJobCmd.MarkFlagRequired("execution-role-arn")
	sagemakerGeospatial_exportVectorEnrichmentJobCmd.MarkFlagRequired("output-config")
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_exportVectorEnrichmentJobCmd)
}
