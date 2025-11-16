package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listModelCopyJobsCmd = &cobra.Command{
	Use:   "list-model-copy-jobs",
	Short: "Returns a list of model copy jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listModelCopyJobsCmd).Standalone()

	bedrock_listModelCopyJobsCmd.Flags().String("creation-time-after", "", "Filters for model copy jobs created after the specified time.")
	bedrock_listModelCopyJobsCmd.Flags().String("creation-time-before", "", "Filters for model copy jobs created before the specified time.")
	bedrock_listModelCopyJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrock_listModelCopyJobsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrock_listModelCopyJobsCmd.Flags().String("sort-by", "", "The field to sort by in the returned list of model copy jobs.")
	bedrock_listModelCopyJobsCmd.Flags().String("sort-order", "", "Specifies whether to sort the results in ascending or descending order.")
	bedrock_listModelCopyJobsCmd.Flags().String("source-account-equals", "", "Filters for model copy jobs in which the account that the source model belongs to is equal to the value that you specify.")
	bedrock_listModelCopyJobsCmd.Flags().String("source-model-arn-equals", "", "Filters for model copy jobs in which the Amazon Resource Name (ARN) of the source model to is equal to the value that you specify.")
	bedrock_listModelCopyJobsCmd.Flags().String("status-equals", "", "Filters for model copy jobs whose status matches the value that you specify.")
	bedrock_listModelCopyJobsCmd.Flags().String("target-model-name-contains", "", "Filters for model copy jobs in which the name of the copied model contains the string that you specify.")
	bedrockCmd.AddCommand(bedrock_listModelCopyJobsCmd)
}
