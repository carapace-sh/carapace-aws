package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_updateTelemetryRuleForOrganizationCmd = &cobra.Command{
	Use:   "update-telemetry-rule-for-organization",
	Short: "Updates an existing telemetry rule that applies across an Amazon Web Services Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_updateTelemetryRuleForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_updateTelemetryRuleForOrganizationCmd).Standalone()

		observabilityadmin_updateTelemetryRuleForOrganizationCmd.Flags().String("rule", "", "The new configuration details for the organization telemetry rule, including resource type, telemetry type, and destination configuration.")
		observabilityadmin_updateTelemetryRuleForOrganizationCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the organization telemetry rule to update.")
		observabilityadmin_updateTelemetryRuleForOrganizationCmd.MarkFlagRequired("rule")
		observabilityadmin_updateTelemetryRuleForOrganizationCmd.MarkFlagRequired("rule-identifier")
	})
	observabilityadminCmd.AddCommand(observabilityadmin_updateTelemetryRuleForOrganizationCmd)
}
