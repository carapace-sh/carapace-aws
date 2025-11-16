package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_startDetectorModelAnalysisCmd = &cobra.Command{
	Use:   "start-detector-model-analysis",
	Short: "Performs an analysis of your detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_startDetectorModelAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_startDetectorModelAnalysisCmd).Standalone()

		iotevents_startDetectorModelAnalysisCmd.Flags().String("detector-model-definition", "", "")
		iotevents_startDetectorModelAnalysisCmd.MarkFlagRequired("detector-model-definition")
	})
	ioteventsCmd.AddCommand(iotevents_startDetectorModelAnalysisCmd)
}
