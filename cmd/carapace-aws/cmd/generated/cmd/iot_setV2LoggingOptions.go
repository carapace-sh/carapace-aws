package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_setV2LoggingOptionsCmd = &cobra.Command{
	Use:   "set-v2-logging-options",
	Short: "Sets the logging options for the V2 logging service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_setV2LoggingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_setV2LoggingOptionsCmd).Standalone()

		iot_setV2LoggingOptionsCmd.Flags().String("default-log-level", "", "The default logging level.")
		iot_setV2LoggingOptionsCmd.Flags().String("disable-all-logs", "", "If true all logs are disabled.")
		iot_setV2LoggingOptionsCmd.Flags().String("role-arn", "", "The ARN of the role that allows IoT to write to Cloudwatch logs.")
	})
	iotCmd.AddCommand(iot_setV2LoggingOptionsCmd)
}
