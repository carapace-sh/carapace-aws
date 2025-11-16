package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_startTelemetryEvaluationForOrganizationCmd = &cobra.Command{
	Use:   "start-telemetry-evaluation-for-organization",
	Short: "This actions begins onboarding the organization and all member accounts to the telemetry config feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_startTelemetryEvaluationForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_startTelemetryEvaluationForOrganizationCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_startTelemetryEvaluationForOrganizationCmd)
}
