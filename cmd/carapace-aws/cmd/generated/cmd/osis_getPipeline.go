package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_getPipelineCmd = &cobra.Command{
	Use:   "get-pipeline",
	Short: "Retrieves information about an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_getPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osis_getPipelineCmd).Standalone()

		osis_getPipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline.")
		osis_getPipelineCmd.MarkFlagRequired("pipeline-name")
	})
	osisCmd.AddCommand(osis_getPipelineCmd)
}
