package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_stopTelemetryEvaluationForOrganizationCmd = &cobra.Command{
	Use:   "stop-telemetry-evaluation-for-organization",
	Short: "This action offboards the Organization of the caller Amazon Web Services account from the telemetry config feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_stopTelemetryEvaluationForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_stopTelemetryEvaluationForOrganizationCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_stopTelemetryEvaluationForOrganizationCmd)
}
