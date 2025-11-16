package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_startTelemetryEnrichmentCmd = &cobra.Command{
	Use:   "start-telemetry-enrichment",
	Short: "Enables the resource tags for telemetry feature for your account, which enhances telemetry data with additional resource metadata from Amazon Web Services Resource Explorer to provide richer context for monitoring and observability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_startTelemetryEnrichmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_startTelemetryEnrichmentCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_startTelemetryEnrichmentCmd)
}
