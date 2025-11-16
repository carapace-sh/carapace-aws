package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listFlowAliasesCmd = &cobra.Command{
	Use:   "list-flow-aliases",
	Short: "Returns a list of aliases for a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listFlowAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listFlowAliasesCmd).Standalone()

		bedrockAgent_listFlowAliasesCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow for which aliases are being returned.")
		bedrockAgent_listFlowAliasesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listFlowAliasesCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listFlowAliasesCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listFlowAliasesCmd)
}
