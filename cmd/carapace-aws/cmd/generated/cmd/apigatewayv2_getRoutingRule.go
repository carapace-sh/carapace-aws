package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getRoutingRuleCmd = &cobra.Command{
	Use:   "get-routing-rule",
	Short: "Gets a routing rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getRoutingRuleCmd).Standalone()

	apigatewayv2_getRoutingRuleCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_getRoutingRuleCmd.Flags().String("domain-name-id", "", "The domain name ID.")
	apigatewayv2_getRoutingRuleCmd.Flags().String("routing-rule-id", "", "The routing rule ID.")
	apigatewayv2_getRoutingRuleCmd.MarkFlagRequired("domain-name")
	apigatewayv2_getRoutingRuleCmd.MarkFlagRequired("routing-rule-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getRoutingRuleCmd)
}
