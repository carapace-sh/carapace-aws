package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_getPipelineChangeProgressCmd = &cobra.Command{
	Use:   "get-pipeline-change-progress",
	Short: "Returns progress information for the current change happening on an OpenSearch Ingestion pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_getPipelineChangeProgressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osis_getPipelineChangeProgressCmd).Standalone()

		osis_getPipelineChangeProgressCmd.Flags().String("pipeline-name", "", "The name of the pipeline.")
		osis_getPipelineChangeProgressCmd.MarkFlagRequired("pipeline-name")
	})
	osisCmd.AddCommand(osis_getPipelineChangeProgressCmd)
}
