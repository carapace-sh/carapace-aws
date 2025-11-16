package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_describeDetectorModelAnalysisCmd = &cobra.Command{
	Use:   "describe-detector-model-analysis",
	Short: "Retrieves runtime information about a detector model analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_describeDetectorModelAnalysisCmd).Standalone()

	iotevents_describeDetectorModelAnalysisCmd.Flags().String("analysis-id", "", "The ID of the analysis result that you want to retrieve.")
	iotevents_describeDetectorModelAnalysisCmd.MarkFlagRequired("analysis-id")
	ioteventsCmd.AddCommand(iotevents_describeDetectorModelAnalysisCmd)
}
