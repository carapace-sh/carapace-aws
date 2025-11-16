package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_describePipelineCmd = &cobra.Command{
	Use:   "describe-pipeline",
	Short: "Retrieves information about a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_describePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_describePipelineCmd).Standalone()

		iotanalytics_describePipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline whose information is retrieved.")
		iotanalytics_describePipelineCmd.MarkFlagRequired("pipeline-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_describePipelineCmd)
}
