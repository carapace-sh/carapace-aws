package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "Retrieves a list of pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_listPipelinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_listPipelinesCmd).Standalone()

		iotanalytics_listPipelinesCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		iotanalytics_listPipelinesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_listPipelinesCmd)
}
