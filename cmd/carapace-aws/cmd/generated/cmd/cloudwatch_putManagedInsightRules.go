package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putManagedInsightRulesCmd = &cobra.Command{
	Use:   "put-managed-insight-rules",
	Short: "Creates a managed Contributor Insights rule for a specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putManagedInsightRulesCmd).Standalone()

	cloudwatch_putManagedInsightRulesCmd.Flags().String("managed-rules", "", "A list of `ManagedRules` to enable.")
	cloudwatch_putManagedInsightRulesCmd.MarkFlagRequired("managed-rules")
	cloudwatchCmd.AddCommand(cloudwatch_putManagedInsightRulesCmd)
}
