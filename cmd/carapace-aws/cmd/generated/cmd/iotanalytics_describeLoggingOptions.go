package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_describeLoggingOptionsCmd = &cobra.Command{
	Use:   "describe-logging-options",
	Short: "Retrieves the current settings of the IoT Analytics logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_describeLoggingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_describeLoggingOptionsCmd).Standalone()

	})
	iotanalyticsCmd.AddCommand(iotanalytics_describeLoggingOptionsCmd)
}
