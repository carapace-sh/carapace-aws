package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listModelCustomizationJobsCmd = &cobra.Command{
	Use:   "list-model-customization-jobs",
	Short: "Returns a list of model customization jobs that you have submitted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listModelCustomizationJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_listModelCustomizationJobsCmd).Standalone()

		bedrock_listModelCustomizationJobsCmd.Flags().String("creation-time-after", "", "Return customization jobs created after the specified time.")
		bedrock_listModelCustomizationJobsCmd.Flags().String("creation-time-before", "", "Return customization jobs created before the specified time.")
		bedrock_listModelCustomizationJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrock_listModelCustomizationJobsCmd.Flags().String("name-contains", "", "Return customization jobs only if the job name contains these characters.")
		bedrock_listModelCustomizationJobsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrock_listModelCustomizationJobsCmd.Flags().String("sort-by", "", "The field to sort by in the returned list of jobs.")
		bedrock_listModelCustomizationJobsCmd.Flags().String("sort-order", "", "The sort order of the results.")
		bedrock_listModelCustomizationJobsCmd.Flags().String("status-equals", "", "Return customization jobs with the specified status.")
	})
	bedrockCmd.AddCommand(bedrock_listModelCustomizationJobsCmd)
}
