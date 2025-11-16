package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listGatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: "Lists all gateways in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listGatewaysCmd).Standalone()

	bedrockAgentcoreControl_listGatewaysCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgentcoreControl_listGatewaysCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listGatewaysCmd)
}
