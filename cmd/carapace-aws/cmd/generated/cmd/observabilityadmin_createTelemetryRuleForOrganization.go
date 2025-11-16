package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_createTelemetryRuleForOrganizationCmd = &cobra.Command{
	Use:   "create-telemetry-rule-for-organization",
	Short: "Creates a telemetry rule that applies across an Amazon Web Services Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_createTelemetryRuleForOrganizationCmd).Standalone()

	observabilityadmin_createTelemetryRuleForOrganizationCmd.Flags().String("rule", "", "The configuration details for the organization-wide telemetry rule, including the resource type, telemetry type, destination configuration, and selection criteria for which resources the rule applies to across the organization.")
	observabilityadmin_createTelemetryRuleForOrganizationCmd.Flags().String("rule-name", "", "A unique name for the organization-wide telemetry rule being created.")
	observabilityadmin_createTelemetryRuleForOrganizationCmd.Flags().String("tags", "", "The key-value pairs to associate with the organization telemetry rule resource for categorization and management purposes.")
	observabilityadmin_createTelemetryRuleForOrganizationCmd.MarkFlagRequired("rule")
	observabilityadmin_createTelemetryRuleForOrganizationCmd.MarkFlagRequired("rule-name")
	observabilityadminCmd.AddCommand(observabilityadmin_createTelemetryRuleForOrganizationCmd)
}
