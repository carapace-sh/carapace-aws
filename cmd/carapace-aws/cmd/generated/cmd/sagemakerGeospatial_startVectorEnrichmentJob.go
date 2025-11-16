package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_startVectorEnrichmentJobCmd = &cobra.Command{
	Use:   "start-vector-enrichment-job",
	Short: "Creates a Vector Enrichment job for the supplied job type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_startVectorEnrichmentJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_startVectorEnrichmentJobCmd).Standalone()

		sagemakerGeospatial_startVectorEnrichmentJobCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that you specified for the job.")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.Flags().String("input-config", "", "Input configuration information for the Vector Enrichment job.")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.Flags().String("job-config", "", "An object containing information about the job configuration.")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.Flags().String("kms-key-id", "", "The Key Management Service key ID for server-side encryption.")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.Flags().String("name", "", "The name of the Vector Enrichment job.")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.Flags().String("tags", "", "Each tag consists of a key and a value.")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.MarkFlagRequired("execution-role-arn")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.MarkFlagRequired("input-config")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.MarkFlagRequired("job-config")
		sagemakerGeospatial_startVectorEnrichmentJobCmd.MarkFlagRequired("name")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_startVectorEnrichmentJobCmd)
}
