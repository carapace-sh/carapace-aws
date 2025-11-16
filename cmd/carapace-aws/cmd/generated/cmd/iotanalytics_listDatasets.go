package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Retrieves information about datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_listDatasetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_listDatasetsCmd).Standalone()

		iotanalytics_listDatasetsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		iotanalytics_listDatasetsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_listDatasetsCmd)
}
