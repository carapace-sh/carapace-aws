package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_updateTelemetryRuleCmd = &cobra.Command{
	Use:   "update-telemetry-rule",
	Short: "Updates an existing telemetry rule in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_updateTelemetryRuleCmd).Standalone()

	observabilityadmin_updateTelemetryRuleCmd.Flags().String("rule", "", "The new configuration details for the telemetry rule.")
	observabilityadmin_updateTelemetryRuleCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the telemetry rule to update.")
	observabilityadmin_updateTelemetryRuleCmd.MarkFlagRequired("rule")
	observabilityadmin_updateTelemetryRuleCmd.MarkFlagRequired("rule-identifier")
	observabilityadminCmd.AddCommand(observabilityadmin_updateTelemetryRuleCmd)
}
