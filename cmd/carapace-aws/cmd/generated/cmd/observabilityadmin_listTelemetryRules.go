package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_listTelemetryRulesCmd = &cobra.Command{
	Use:   "list-telemetry-rules",
	Short: "Lists all telemetry rules in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_listTelemetryRulesCmd).Standalone()

	observabilityadmin_listTelemetryRulesCmd.Flags().String("max-results", "", "The maximum number of telemetry rules to return in a single call.")
	observabilityadmin_listTelemetryRulesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	observabilityadmin_listTelemetryRulesCmd.Flags().String("rule-name-prefix", "", "A string to filter telemetry rules whose names begin with the specified prefix.")
	observabilityadminCmd.AddCommand(observabilityadmin_listTelemetryRulesCmd)
}
