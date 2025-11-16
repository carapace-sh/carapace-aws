package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_listDetectorModelsCmd = &cobra.Command{
	Use:   "list-detector-models",
	Short: "Lists the detector models you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_listDetectorModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_listDetectorModelsCmd).Standalone()

		iotevents_listDetectorModelsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		iotevents_listDetectorModelsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
	})
	ioteventsCmd.AddCommand(iotevents_listDetectorModelsCmd)
}
