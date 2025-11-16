package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_putRoutingRuleCmd = &cobra.Command{
	Use:   "put-routing-rule",
	Short: "Updates a routing rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_putRoutingRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_putRoutingRuleCmd).Standalone()

		apigatewayv2_putRoutingRuleCmd.Flags().String("actions", "", "The routing rule action.")
		apigatewayv2_putRoutingRuleCmd.Flags().String("conditions", "", "The routing rule condition.")
		apigatewayv2_putRoutingRuleCmd.Flags().String("domain-name", "", "The domain name.")
		apigatewayv2_putRoutingRuleCmd.Flags().String("domain-name-id", "", "The domain name ID.")
		apigatewayv2_putRoutingRuleCmd.Flags().String("priority", "", "The routing rule priority.")
		apigatewayv2_putRoutingRuleCmd.Flags().String("routing-rule-id", "", "The routing rule ID.")
		apigatewayv2_putRoutingRuleCmd.MarkFlagRequired("actions")
		apigatewayv2_putRoutingRuleCmd.MarkFlagRequired("conditions")
		apigatewayv2_putRoutingRuleCmd.MarkFlagRequired("domain-name")
		apigatewayv2_putRoutingRuleCmd.MarkFlagRequired("priority")
		apigatewayv2_putRoutingRuleCmd.MarkFlagRequired("routing-rule-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_putRoutingRuleCmd)
}
