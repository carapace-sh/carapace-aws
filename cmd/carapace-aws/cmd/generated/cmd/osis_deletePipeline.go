package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "Deletes an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_deletePipelineCmd).Standalone()

	osis_deletePipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to delete.")
	osis_deletePipelineCmd.MarkFlagRequired("pipeline-name")
	osisCmd.AddCommand(osis_deletePipelineCmd)
}
