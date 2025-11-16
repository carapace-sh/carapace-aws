package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteRoutingRuleCmd = &cobra.Command{
	Use:   "delete-routing-rule",
	Short: "Deletes a routing rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteRoutingRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_deleteRoutingRuleCmd).Standalone()

		apigatewayv2_deleteRoutingRuleCmd.Flags().String("domain-name", "", "The domain name.")
		apigatewayv2_deleteRoutingRuleCmd.Flags().String("domain-name-id", "", "The domain name ID.")
		apigatewayv2_deleteRoutingRuleCmd.Flags().String("routing-rule-id", "", "The routing rule ID.")
		apigatewayv2_deleteRoutingRuleCmd.MarkFlagRequired("domain-name")
		apigatewayv2_deleteRoutingRuleCmd.MarkFlagRequired("routing-rule-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteRoutingRuleCmd)
}
