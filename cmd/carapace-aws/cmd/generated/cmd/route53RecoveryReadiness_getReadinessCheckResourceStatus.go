package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_getReadinessCheckResourceStatusCmd = &cobra.Command{
	Use:   "get-readiness-check-resource-status",
	Short: "Gets individual readiness status for a readiness check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_getReadinessCheckResourceStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53RecoveryReadiness_getReadinessCheckResourceStatusCmd).Standalone()

		route53RecoveryReadiness_getReadinessCheckResourceStatusCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		route53RecoveryReadiness_getReadinessCheckResourceStatusCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
		route53RecoveryReadiness_getReadinessCheckResourceStatusCmd.Flags().String("readiness-check-name", "", "Name of a readiness check.")
		route53RecoveryReadiness_getReadinessCheckResourceStatusCmd.Flags().String("resource-identifier", "", "The resource identifier, which is the Amazon Resource Name (ARN) or the identifier generated for the resource by Application Recovery Controller (for example, for a DNS target resource).")
		route53RecoveryReadiness_getReadinessCheckResourceStatusCmd.MarkFlagRequired("readiness-check-name")
		route53RecoveryReadiness_getReadinessCheckResourceStatusCmd.MarkFlagRequired("resource-identifier")
	})
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_getReadinessCheckResourceStatusCmd)
}
