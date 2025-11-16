package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_updatePipelineCmd = &cobra.Command{
	Use:   "update-pipeline",
	Short: "Updates an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_updatePipelineCmd).Standalone()

	osis_updatePipelineCmd.Flags().String("buffer-options", "", "Key-value pairs to configure persistent buffering for the pipeline.")
	osis_updatePipelineCmd.Flags().String("encryption-at-rest-options", "", "Key-value pairs to configure encryption for data that is written to a persistent buffer.")
	osis_updatePipelineCmd.Flags().String("log-publishing-options", "", "Key-value pairs to configure log publishing.")
	osis_updatePipelineCmd.Flags().String("max-units", "", "The maximum pipeline capacity, in Ingestion Compute Units (ICUs)")
	osis_updatePipelineCmd.Flags().String("min-units", "", "The minimum pipeline capacity, in Ingestion Compute Units (ICUs).")
	osis_updatePipelineCmd.Flags().String("pipeline-configuration-body", "", "The pipeline configuration in YAML format.")
	osis_updatePipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to update.")
	osis_updatePipelineCmd.Flags().String("pipeline-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants the pipeline permission to access Amazon Web Services resources.")
	osis_updatePipelineCmd.MarkFlagRequired("pipeline-name")
	osisCmd.AddCommand(osis_updatePipelineCmd)
}
