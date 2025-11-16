package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53RecoveryReadiness_listRulesCmd = &cobra.Command{
	Use:   "list-rules",
	Short: "Lists all readiness rules, or lists the readiness rules for a specific resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53RecoveryReadiness_listRulesCmd).Standalone()

	route53RecoveryReadiness_listRulesCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	route53RecoveryReadiness_listRulesCmd.Flags().String("next-token", "", "The token that identifies which batch of results you want to see.")
	route53RecoveryReadiness_listRulesCmd.Flags().String("resource-type", "", "The resource type that a readiness rule applies to.")
	route53RecoveryReadinessCmd.AddCommand(route53RecoveryReadiness_listRulesCmd)
}
