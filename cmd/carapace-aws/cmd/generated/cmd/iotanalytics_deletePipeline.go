package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "Deletes the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_deletePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_deletePipelineCmd).Standalone()

		iotanalytics_deletePipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to delete.")
		iotanalytics_deletePipelineCmd.MarkFlagRequired("pipeline-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_deletePipelineCmd)
}
