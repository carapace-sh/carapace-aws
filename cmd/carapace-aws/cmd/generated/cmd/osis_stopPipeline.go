package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_stopPipelineCmd = &cobra.Command{
	Use:   "stop-pipeline",
	Short: "Stops an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_stopPipelineCmd).Standalone()

	osis_stopPipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to stop.")
	osis_stopPipelineCmd.MarkFlagRequired("pipeline-name")
	osisCmd.AddCommand(osis_stopPipelineCmd)
}
