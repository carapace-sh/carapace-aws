package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_getTelemetryEvaluationStatusCmd = &cobra.Command{
	Use:   "get-telemetry-evaluation-status",
	Short: "Returns the current onboarding status of the telemetry config feature, including the status of the feature and reason the feature failed to start or stop.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_getTelemetryEvaluationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_getTelemetryEvaluationStatusCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_getTelemetryEvaluationStatusCmd)
}
