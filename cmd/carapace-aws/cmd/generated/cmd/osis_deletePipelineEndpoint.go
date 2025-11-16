package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_deletePipelineEndpointCmd = &cobra.Command{
	Use:   "delete-pipeline-endpoint",
	Short: "Deletes a VPC endpoint for an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_deletePipelineEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osis_deletePipelineEndpointCmd).Standalone()

		osis_deletePipelineEndpointCmd.Flags().String("endpoint-id", "", "The unique identifier of the pipeline endpoint to delete.")
		osis_deletePipelineEndpointCmd.MarkFlagRequired("endpoint-id")
	})
	osisCmd.AddCommand(osis_deletePipelineEndpointCmd)
}
