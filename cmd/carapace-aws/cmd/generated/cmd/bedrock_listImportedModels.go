package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listImportedModelsCmd = &cobra.Command{
	Use:   "list-imported-models",
	Short: "Returns a list of models you've imported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listImportedModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_listImportedModelsCmd).Standalone()

		bedrock_listImportedModelsCmd.Flags().String("creation-time-after", "", "Return imported models that were created after the specified time.")
		bedrock_listImportedModelsCmd.Flags().String("creation-time-before", "", "Return imported models that created before the specified time.")
		bedrock_listImportedModelsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrock_listImportedModelsCmd.Flags().String("name-contains", "", "Return imported models only if the model name contains these characters.")
		bedrock_listImportedModelsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrock_listImportedModelsCmd.Flags().String("sort-by", "", "The field to sort by in the returned list of imported models.")
		bedrock_listImportedModelsCmd.Flags().String("sort-order", "", "Specifies whetehr to sort the results in ascending or descending order.")
	})
	bedrockCmd.AddCommand(bedrock_listImportedModelsCmd)
}
