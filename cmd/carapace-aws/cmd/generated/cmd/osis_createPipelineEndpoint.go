package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_createPipelineEndpointCmd = &cobra.Command{
	Use:   "create-pipeline-endpoint",
	Short: "Creates a VPC endpoint for an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_createPipelineEndpointCmd).Standalone()

	osis_createPipelineEndpointCmd.Flags().String("pipeline-arn", "", "The Amazon Resource Name (ARN) of the pipeline to create the endpoint for.")
	osis_createPipelineEndpointCmd.Flags().String("vpc-options", "", "Container for the VPC configuration for the pipeline endpoint, including subnet IDs and security group IDs.")
	osis_createPipelineEndpointCmd.MarkFlagRequired("pipeline-arn")
	osis_createPipelineEndpointCmd.MarkFlagRequired("vpc-options")
	osisCmd.AddCommand(osis_createPipelineEndpointCmd)
}
