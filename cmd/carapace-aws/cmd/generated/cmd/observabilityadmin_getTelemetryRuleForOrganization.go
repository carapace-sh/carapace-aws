package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_getTelemetryRuleForOrganizationCmd = &cobra.Command{
	Use:   "get-telemetry-rule-for-organization",
	Short: "Retrieves the details of a specific organization telemetry rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_getTelemetryRuleForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_getTelemetryRuleForOrganizationCmd).Standalone()

		observabilityadmin_getTelemetryRuleForOrganizationCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the organization telemetry rule to retrieve.")
		observabilityadmin_getTelemetryRuleForOrganizationCmd.MarkFlagRequired("rule-identifier")
	})
	observabilityadminCmd.AddCommand(observabilityadmin_getTelemetryRuleForOrganizationCmd)
}
