package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listGatewayTargetsCmd = &cobra.Command{
	Use:   "list-gateway-targets",
	Short: "Lists all targets for a specific gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listGatewayTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_listGatewayTargetsCmd).Standalone()

		bedrockAgentcoreControl_listGatewayTargetsCmd.Flags().String("gateway-identifier", "", "The identifier of the gateway to list targets for.")
		bedrockAgentcoreControl_listGatewayTargetsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgentcoreControl_listGatewayTargetsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgentcoreControl_listGatewayTargetsCmd.MarkFlagRequired("gateway-identifier")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listGatewayTargetsCmd)
}
