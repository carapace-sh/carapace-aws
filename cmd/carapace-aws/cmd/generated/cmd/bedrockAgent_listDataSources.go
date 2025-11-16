package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listDataSourcesCmd = &cobra.Command{
	Use:   "list-data-sources",
	Short: "Lists the data sources in a knowledge base and information about each one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listDataSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listDataSourcesCmd).Standalone()

		bedrockAgent_listDataSourcesCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base for which to return a list of information.")
		bedrockAgent_listDataSourcesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listDataSourcesCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listDataSourcesCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listDataSourcesCmd)
}
