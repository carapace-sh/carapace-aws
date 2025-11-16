package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listModelImportJobsCmd = &cobra.Command{
	Use:   "list-model-import-jobs",
	Short: "Returns a list of import jobs you've submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listModelImportJobsCmd).Standalone()

	bedrock_listModelImportJobsCmd.Flags().String("creation-time-after", "", "Return import jobs that were created after the specified time.")
	bedrock_listModelImportJobsCmd.Flags().String("creation-time-before", "", "Return import jobs that were created before the specified time.")
	bedrock_listModelImportJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrock_listModelImportJobsCmd.Flags().String("name-contains", "", "Return imported jobs only if the job name contains these characters.")
	bedrock_listModelImportJobsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrock_listModelImportJobsCmd.Flags().String("sort-by", "", "The field to sort by in the returned list of imported jobs.")
	bedrock_listModelImportJobsCmd.Flags().String("sort-order", "", "Specifies whether to sort the results in ascending or descending order.")
	bedrock_listModelImportJobsCmd.Flags().String("status-equals", "", "Return imported jobs with the specified status.")
	bedrockCmd.AddCommand(bedrock_listModelImportJobsCmd)
}
