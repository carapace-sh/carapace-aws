package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_stopTelemetryEvaluationCmd = &cobra.Command{
	Use:   "stop-telemetry-evaluation",
	Short: "This action begins offboarding the caller Amazon Web Services account from the telemetry config feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_stopTelemetryEvaluationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_stopTelemetryEvaluationCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_stopTelemetryEvaluationCmd)
}
