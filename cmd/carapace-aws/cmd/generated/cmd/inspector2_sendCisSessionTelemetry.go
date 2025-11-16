package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_sendCisSessionTelemetryCmd = &cobra.Command{
	Use:   "send-cis-session-telemetry",
	Short: "Sends a CIS session telemetry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_sendCisSessionTelemetryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_sendCisSessionTelemetryCmd).Standalone()

		inspector2_sendCisSessionTelemetryCmd.Flags().String("messages", "", "The CIS session telemetry messages.")
		inspector2_sendCisSessionTelemetryCmd.Flags().String("scan-job-id", "", "A unique identifier for the scan job.")
		inspector2_sendCisSessionTelemetryCmd.Flags().String("session-token", "", "The unique token that identifies the CIS session.")
		inspector2_sendCisSessionTelemetryCmd.MarkFlagRequired("messages")
		inspector2_sendCisSessionTelemetryCmd.MarkFlagRequired("scan-job-id")
		inspector2_sendCisSessionTelemetryCmd.MarkFlagRequired("session-token")
	})
	inspector2Cmd.AddCommand(inspector2_sendCisSessionTelemetryCmd)
}
