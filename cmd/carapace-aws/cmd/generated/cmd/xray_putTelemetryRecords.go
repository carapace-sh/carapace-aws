package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_putTelemetryRecordsCmd = &cobra.Command{
	Use:   "put-telemetry-records",
	Short: "Used by the Amazon Web Services X-Ray daemon to upload telemetry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_putTelemetryRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_putTelemetryRecordsCmd).Standalone()

		xray_putTelemetryRecordsCmd.Flags().String("ec2-instance-id", "", "")
		xray_putTelemetryRecordsCmd.Flags().String("hostname", "", "")
		xray_putTelemetryRecordsCmd.Flags().String("resource-arn", "", "")
		xray_putTelemetryRecordsCmd.Flags().String("telemetry-records", "", "")
		xray_putTelemetryRecordsCmd.MarkFlagRequired("telemetry-records")
	})
	xrayCmd.AddCommand(xray_putTelemetryRecordsCmd)
}
