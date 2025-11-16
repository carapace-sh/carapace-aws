package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_listDatasetContentsCmd = &cobra.Command{
	Use:   "list-dataset-contents",
	Short: "Lists information about dataset contents that have been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_listDatasetContentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_listDatasetContentsCmd).Standalone()

		iotanalytics_listDatasetContentsCmd.Flags().String("dataset-name", "", "The name of the dataset whose contents information you want to list.")
		iotanalytics_listDatasetContentsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		iotanalytics_listDatasetContentsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		iotanalytics_listDatasetContentsCmd.Flags().String("scheduled-before", "", "A filter to limit results to those dataset contents whose creation is scheduled before the given time.")
		iotanalytics_listDatasetContentsCmd.Flags().String("scheduled-on-or-after", "", "A filter to limit results to those dataset contents whose creation is scheduled on or after the given time.")
		iotanalytics_listDatasetContentsCmd.MarkFlagRequired("dataset-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_listDatasetContentsCmd)
}
