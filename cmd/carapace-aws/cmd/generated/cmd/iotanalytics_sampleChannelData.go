package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_sampleChannelDataCmd = &cobra.Command{
	Use:   "sample-channel-data",
	Short: "Retrieves a sample of messages from the specified channel ingested during the specified timeframe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_sampleChannelDataCmd).Standalone()

	iotanalytics_sampleChannelDataCmd.Flags().String("channel-name", "", "The name of the channel whose message samples are retrieved.")
	iotanalytics_sampleChannelDataCmd.Flags().String("end-time", "", "The end of the time window from which sample messages are retrieved.")
	iotanalytics_sampleChannelDataCmd.Flags().String("max-messages", "", "The number of sample messages to be retrieved.")
	iotanalytics_sampleChannelDataCmd.Flags().String("start-time", "", "The start of the time window from which sample messages are retrieved.")
	iotanalytics_sampleChannelDataCmd.MarkFlagRequired("channel-name")
	iotanalyticsCmd.AddCommand(iotanalytics_sampleChannelDataCmd)
}
