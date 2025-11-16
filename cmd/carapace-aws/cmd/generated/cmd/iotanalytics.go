package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalyticsCmd = &cobra.Command{
	Use:   "iotanalytics",
	Short: "IoT Analytics allows you to collect large amounts of device data, process messages, and store them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalyticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalyticsCmd).Standalone()

	})
	rootCmd.AddCommand(iotanalyticsCmd)
}
