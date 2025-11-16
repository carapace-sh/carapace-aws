package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createRoutingRuleCmd = &cobra.Command{
	Use:   "create-routing-rule",
	Short: "Creates a RoutingRule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createRoutingRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_createRoutingRuleCmd).Standalone()

		apigatewayv2_createRoutingRuleCmd.Flags().String("actions", "", "Represents a routing rule action.")
		apigatewayv2_createRoutingRuleCmd.Flags().String("conditions", "", "Represents a condition.")
		apigatewayv2_createRoutingRuleCmd.Flags().String("domain-name", "", "The domain name.")
		apigatewayv2_createRoutingRuleCmd.Flags().String("domain-name-id", "", "The domain name ID.")
		apigatewayv2_createRoutingRuleCmd.Flags().String("priority", "", "Represents the priority of the routing rule.")
		apigatewayv2_createRoutingRuleCmd.MarkFlagRequired("actions")
		apigatewayv2_createRoutingRuleCmd.MarkFlagRequired("conditions")
		apigatewayv2_createRoutingRuleCmd.MarkFlagRequired("domain-name")
		apigatewayv2_createRoutingRuleCmd.MarkFlagRequired("priority")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_createRoutingRuleCmd)
}
