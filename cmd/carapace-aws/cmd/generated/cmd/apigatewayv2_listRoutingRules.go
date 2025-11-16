package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_listRoutingRulesCmd = &cobra.Command{
	Use:   "list-routing-rules",
	Short: "Lists routing rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_listRoutingRulesCmd).Standalone()

	apigatewayv2_listRoutingRulesCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_listRoutingRulesCmd.Flags().String("domain-name-id", "", "The domain name ID.")
	apigatewayv2_listRoutingRulesCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
	apigatewayv2_listRoutingRulesCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
	apigatewayv2_listRoutingRulesCmd.MarkFlagRequired("domain-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_listRoutingRulesCmd)
}
