package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_deleteTelemetryRuleForOrganizationCmd = &cobra.Command{
	Use:   "delete-telemetry-rule-for-organization",
	Short: "Deletes an organization-wide telemetry rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_deleteTelemetryRuleForOrganizationCmd).Standalone()

	observabilityadmin_deleteTelemetryRuleForOrganizationCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the organization telemetry rule to delete.")
	observabilityadmin_deleteTelemetryRuleForOrganizationCmd.MarkFlagRequired("rule-identifier")
	observabilityadminCmd.AddCommand(observabilityadmin_deleteTelemetryRuleForOrganizationCmd)
}
