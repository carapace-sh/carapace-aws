package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listFlowVersionsCmd = &cobra.Command{
	Use:   "list-flow-versions",
	Short: "Returns a list of information about each flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listFlowVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listFlowVersionsCmd).Standalone()

		bedrockAgent_listFlowVersionsCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgent_listFlowVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listFlowVersionsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listFlowVersionsCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listFlowVersionsCmd)
}
