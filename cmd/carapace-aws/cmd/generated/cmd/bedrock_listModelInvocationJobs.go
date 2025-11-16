package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listModelInvocationJobsCmd = &cobra.Command{
	Use:   "list-model-invocation-jobs",
	Short: "Lists all batch inference jobs in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listModelInvocationJobsCmd).Standalone()

	bedrock_listModelInvocationJobsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	bedrock_listModelInvocationJobsCmd.Flags().String("name-contains", "", "Specify a string to filter for batch inference jobs whose names contain the string.")
	bedrock_listModelInvocationJobsCmd.Flags().String("next-token", "", "If there were more results than the value you specified in the `maxResults` field in a previous `ListModelInvocationJobs` request, the response would have returned a `nextToken` value.")
	bedrock_listModelInvocationJobsCmd.Flags().String("sort-by", "", "An attribute by which to sort the results.")
	bedrock_listModelInvocationJobsCmd.Flags().String("sort-order", "", "Specifies whether to sort the results by ascending or descending order.")
	bedrock_listModelInvocationJobsCmd.Flags().String("status-equals", "", "Specify a status to filter for batch inference jobs whose statuses match the string you specify.")
	bedrock_listModelInvocationJobsCmd.Flags().String("submit-time-after", "", "Specify a time to filter for batch inference jobs that were submitted after the time you specify.")
	bedrock_listModelInvocationJobsCmd.Flags().String("submit-time-before", "", "Specify a time to filter for batch inference jobs that were submitted before the time you specify.")
	bedrockCmd.AddCommand(bedrock_listModelInvocationJobsCmd)
}
