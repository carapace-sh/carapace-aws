package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_getDetectorModelAnalysisResultsCmd = &cobra.Command{
	Use:   "get-detector-model-analysis-results",
	Short: "Retrieves one or more analysis results of the detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_getDetectorModelAnalysisResultsCmd).Standalone()

	iotevents_getDetectorModelAnalysisResultsCmd.Flags().String("analysis-id", "", "The ID of the analysis result that you want to retrieve.")
	iotevents_getDetectorModelAnalysisResultsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	iotevents_getDetectorModelAnalysisResultsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
	iotevents_getDetectorModelAnalysisResultsCmd.MarkFlagRequired("analysis-id")
	ioteventsCmd.AddCommand(iotevents_getDetectorModelAnalysisResultsCmd)
}
