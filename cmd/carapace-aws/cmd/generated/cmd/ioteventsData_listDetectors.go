package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_listDetectorsCmd = &cobra.Command{
	Use:   "list-detectors",
	Short: "Lists detectors (the instances of a detector model).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_listDetectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_listDetectorsCmd).Standalone()

		ioteventsData_listDetectorsCmd.Flags().String("detector-model-name", "", "The name of the detector model whose detectors (instances) are listed.")
		ioteventsData_listDetectorsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		ioteventsData_listDetectorsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
		ioteventsData_listDetectorsCmd.Flags().String("state-name", "", "A filter that limits results to those detectors (instances) in the given state.")
		ioteventsData_listDetectorsCmd.MarkFlagRequired("detector-model-name")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_listDetectorsCmd)
}
