package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_describeChannelCmd = &cobra.Command{
	Use:   "describe-channel",
	Short: "Retrieves information about a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_describeChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_describeChannelCmd).Standalone()

		iotanalytics_describeChannelCmd.Flags().String("channel-name", "", "The name of the channel whose information is retrieved.")
		iotanalytics_describeChannelCmd.Flags().String("include-statistics", "", "If true, additional statistical information about the channel is included in the response.")
		iotanalytics_describeChannelCmd.MarkFlagRequired("channel-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_describeChannelCmd)
}
