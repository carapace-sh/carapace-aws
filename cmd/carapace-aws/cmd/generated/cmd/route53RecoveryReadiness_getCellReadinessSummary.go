package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getCellReadinessSummaryCmd = &cobra.Command{
	Use:   "get-cell-readiness-summary",
	Short: "Gets readiness for a cell.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getCellReadinessSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_getCellReadinessSummaryCmd).Standalone()

		route53RecoveryReadiness_getCellReadinessSummaryCmd.Flags().String("cell-name", "", "The name of the cell.")
		route53RecoveryReadiness_getCellReadinessSummaryCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryReadiness_getCellReadinessSummaryCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
		route53RecoveryReadiness_getCellReadinessSummaryCmd.MarkFlagRequired("cell-name")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getCellReadinessSummaryCmd)
}
