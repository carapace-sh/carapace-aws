package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_startTelemetryEvaluationCmd = &cobra.Command{
	Use:   "start-telemetry-evaluation",
	Short: "This action begins onboarding the caller Amazon Web Services account to the telemetry config feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_startTelemetryEvaluationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_startTelemetryEvaluationCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_startTelemetryEvaluationCmd)
}
