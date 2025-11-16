package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Creates an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_createPipelineCmd).Standalone()

	osis_createPipelineCmd.Flags().String("buffer-options", "", "Key-value pairs to configure persistent buffering for the pipeline.")
	osis_createPipelineCmd.Flags().String("encryption-at-rest-options", "", "Key-value pairs to configure encryption for data that is written to a persistent buffer.")
	osis_createPipelineCmd.Flags().String("log-publishing-options", "", "Key-value pairs to configure log publishing.")
	osis_createPipelineCmd.Flags().String("max-units", "", "The maximum pipeline capacity, in Ingestion Compute Units (ICUs).")
	osis_createPipelineCmd.Flags().String("min-units", "", "The minimum pipeline capacity, in Ingestion Compute Units (ICUs).")
	osis_createPipelineCmd.Flags().String("pipeline-configuration-body", "", "The pipeline configuration in YAML format.")
	osis_createPipelineCmd.Flags().String("pipeline-name", "", "The name of the OpenSearch Ingestion pipeline to create.")
	osis_createPipelineCmd.Flags().String("pipeline-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants the pipeline permission to access Amazon Web Services resources.")
	osis_createPipelineCmd.Flags().String("tags", "", "List of tags to add to the pipeline upon creation.")
	osis_createPipelineCmd.Flags().String("vpc-options", "", "Container for the values required to configure VPC access for the pipeline.")
	osis_createPipelineCmd.MarkFlagRequired("max-units")
	osis_createPipelineCmd.MarkFlagRequired("min-units")
	osis_createPipelineCmd.MarkFlagRequired("pipeline-configuration-body")
	osis_createPipelineCmd.MarkFlagRequired("pipeline-name")
	osisCmd.AddCommand(osis_createPipelineCmd)
}
