package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_getTelemetryRuleCmd = &cobra.Command{
	Use:   "get-telemetry-rule",
	Short: "Retrieves the details of a specific telemetry rule in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_getTelemetryRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_getTelemetryRuleCmd).Standalone()

		observabilityadmin_getTelemetryRuleCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the telemetry rule to retrieve.")
		observabilityadmin_getTelemetryRuleCmd.MarkFlagRequired("rule-identifier")
	})
	observabilityadminCmd.AddCommand(observabilityadmin_getTelemetryRuleCmd)
}
