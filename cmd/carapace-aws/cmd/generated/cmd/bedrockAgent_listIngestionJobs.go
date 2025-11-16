package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listIngestionJobsCmd = &cobra.Command{
	Use:   "list-ingestion-jobs",
	Short: "Lists the data ingestion jobs for a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listIngestionJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listIngestionJobsCmd).Standalone()

		bedrockAgent_listIngestionJobsCmd.Flags().String("data-source-id", "", "The unique identifier of the data source for the list of data ingestion jobs.")
		bedrockAgent_listIngestionJobsCmd.Flags().String("filters", "", "Contains information about the filters for filtering the data.")
		bedrockAgent_listIngestionJobsCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base for the list of data ingestion jobs.")
		bedrockAgent_listIngestionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listIngestionJobsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listIngestionJobsCmd.Flags().String("sort-by", "", "Contains details about how to sort the data.")
		bedrockAgent_listIngestionJobsCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_listIngestionJobsCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listIngestionJobsCmd)
}
