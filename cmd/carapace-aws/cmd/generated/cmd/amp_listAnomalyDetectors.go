package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_listAnomalyDetectorsCmd = &cobra.Command{
	Use:   "list-anomaly-detectors",
	Short: "Returns a paginated list of anomaly detectors for a workspace with optional filtering by alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_listAnomalyDetectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_listAnomalyDetectorsCmd).Standalone()

		amp_listAnomalyDetectorsCmd.Flags().String("alias", "", "Filters the results to anomaly detectors with the specified alias.")
		amp_listAnomalyDetectorsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		amp_listAnomalyDetectorsCmd.Flags().String("next-token", "", "The pagination token to continue retrieving results.")
		amp_listAnomalyDetectorsCmd.Flags().String("workspace-id", "", "The identifier of the workspace containing the anomaly detectors to list.")
		amp_listAnomalyDetectorsCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_listAnomalyDetectorsCmd)
}
