package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_createTelemetryRuleCmd = &cobra.Command{
	Use:   "create-telemetry-rule",
	Short: "Creates a telemetry rule that defines how telemetry should be configured for Amazon Web Services resources in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_createTelemetryRuleCmd).Standalone()

	observabilityadmin_createTelemetryRuleCmd.Flags().String("rule", "", "The configuration details for the telemetry rule, including the resource type, telemetry type, destination configuration, and selection criteria for which resources the rule applies to.")
	observabilityadmin_createTelemetryRuleCmd.Flags().String("rule-name", "", "A unique name for the telemetry rule being created.")
	observabilityadmin_createTelemetryRuleCmd.Flags().String("tags", "", "The key-value pairs to associate with the telemetry rule resource for categorization and management purposes.")
	observabilityadmin_createTelemetryRuleCmd.MarkFlagRequired("rule")
	observabilityadmin_createTelemetryRuleCmd.MarkFlagRequired("rule-name")
	observabilityadminCmd.AddCommand(observabilityadmin_createTelemetryRuleCmd)
}
