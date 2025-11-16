package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_stopTelemetryEnrichmentCmd = &cobra.Command{
	Use:   "stop-telemetry-enrichment",
	Short: "Disables the resource tags for telemetry feature for your account, stopping the enhancement of telemetry data with additional resource metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_stopTelemetryEnrichmentCmd).Standalone()

	observabilityadminCmd.AddCommand(observabilityadmin_stopTelemetryEnrichmentCmd)
}
