package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_deleteTelemetryRuleCmd = &cobra.Command{
	Use:   "delete-telemetry-rule",
	Short: "Deletes a telemetry rule from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_deleteTelemetryRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_deleteTelemetryRuleCmd).Standalone()

		observabilityadmin_deleteTelemetryRuleCmd.Flags().String("rule-identifier", "", "The identifier (name or ARN) of the telemetry rule to delete.")
		observabilityadmin_deleteTelemetryRuleCmd.MarkFlagRequired("rule-identifier")
	})
	observabilityadminCmd.AddCommand(observabilityadmin_deleteTelemetryRuleCmd)
}
