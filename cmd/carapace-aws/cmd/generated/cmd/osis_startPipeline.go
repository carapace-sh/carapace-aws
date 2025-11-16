package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_startPipelineCmd = &cobra.Command{
	Use:   "start-pipeline",
	Short: "Starts an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_startPipelineCmd).Standalone()

	osis_startPipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to start.")
	osis_startPipelineCmd.MarkFlagRequired("pipeline-name")
	osisCmd.AddCommand(osis_startPipelineCmd)
}
