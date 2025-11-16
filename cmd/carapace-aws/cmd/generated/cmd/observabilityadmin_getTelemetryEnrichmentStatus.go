package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_getTelemetryEnrichmentStatusCmd = &cobra.Command{
	Use:   "get-telemetry-enrichment-status",
	Short: "Returns the current status of the resource tags for telemetry feature, which enhances telemetry data with additional resource metadata from Amazon Web Services Resource Explorer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_getTelemetryEnrichmentStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_getTelemetryEnrichmentStatusCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_getTelemetryEnrichmentStatusCmd)
}
