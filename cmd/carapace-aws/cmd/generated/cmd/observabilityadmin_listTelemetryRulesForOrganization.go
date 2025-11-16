package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_listTelemetryRulesForOrganizationCmd = &cobra.Command{
	Use:   "list-telemetry-rules-for-organization",
	Short: "Lists all telemetry rules in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_listTelemetryRulesForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_listTelemetryRulesForOrganizationCmd).Standalone()

		observabilityadmin_listTelemetryRulesForOrganizationCmd.Flags().String("max-results", "", "The maximum number of organization telemetry rules to return in a single call.")
		observabilityadmin_listTelemetryRulesForOrganizationCmd.Flags().String("next-token", "", "The token for the next set of results.")
		observabilityadmin_listTelemetryRulesForOrganizationCmd.Flags().String("rule-name-prefix", "", "A string to filter organization telemetry rules whose names begin with the specified prefix.")
		observabilityadmin_listTelemetryRulesForOrganizationCmd.Flags().String("source-account-ids", "", "The list of account IDs to filter organization telemetry rules by their source accounts.")
		observabilityadmin_listTelemetryRulesForOrganizationCmd.Flags().String("source-organization-unit-ids", "", "The list of organizational unit IDs to filter organization telemetry rules by their source organizational units.")
	})
	observabilityadminCmd.AddCommand(observabilityadmin_listTelemetryRulesForOrganizationCmd)
}
