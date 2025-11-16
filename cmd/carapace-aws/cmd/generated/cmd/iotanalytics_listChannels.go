package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_listChannelsCmd = &cobra.Command{
	Use:   "list-channels",
	Short: "Retrieves a list of channels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_listChannelsCmd).Standalone()

	iotanalytics_listChannelsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	iotanalytics_listChannelsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iotanalyticsCmd.AddCommand(iotanalytics_listChannelsCmd)
}
