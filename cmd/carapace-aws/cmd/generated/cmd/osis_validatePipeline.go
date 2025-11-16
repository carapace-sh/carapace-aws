package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_validatePipelineCmd = &cobra.Command{
	Use:   "validate-pipeline",
	Short: "Checks whether an OpenSearch Ingestion pipeline configuration is valid prior to creation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_validatePipelineCmd).Standalone()

	osis_validatePipelineCmd.Flags().String("pipeline-configuration-body", "", "The pipeline configuration in YAML format.")
	osis_validatePipelineCmd.MarkFlagRequired("pipeline-configuration-body")
	osisCmd.AddCommand(osis_validatePipelineCmd)
}
