package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_listDetectorModelVersionsCmd = &cobra.Command{
	Use:   "list-detector-model-versions",
	Short: "Lists all the versions of a detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_listDetectorModelVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_listDetectorModelVersionsCmd).Standalone()

		iotevents_listDetectorModelVersionsCmd.Flags().String("detector-model-name", "", "The name of the detector model whose versions are returned.")
		iotevents_listDetectorModelVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		iotevents_listDetectorModelVersionsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
		iotevents_listDetectorModelVersionsCmd.MarkFlagRequired("detector-model-name")
	})
	ioteventsCmd.AddCommand(iotevents_listDetectorModelVersionsCmd)
}
