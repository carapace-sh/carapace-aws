package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_getTelemetryEvaluationStatusForOrganizationCmd = &cobra.Command{
	Use:   "get-telemetry-evaluation-status-for-organization",
	Short: "This returns the onboarding status of the telemetry configuration feature for the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_getTelemetryEvaluationStatusForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_getTelemetryEvaluationStatusForOrganizationCmd).Standalone()

	})
	observabilityadminCmd.AddCommand(observabilityadmin_getTelemetryEvaluationStatusForOrganizationCmd)
}
