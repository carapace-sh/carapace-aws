package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listCustomModelsCmd = &cobra.Command{
	Use:   "list-custom-models",
	Short: "Returns a list of the custom models that you have created with the `CreateModelCustomizationJob` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listCustomModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_listCustomModelsCmd).Standalone()

		bedrock_listCustomModelsCmd.Flags().String("base-model-arn-equals", "", "Return custom models only if the base model Amazon Resource Name (ARN) matches this parameter.")
		bedrock_listCustomModelsCmd.Flags().String("creation-time-after", "", "Return custom models created after the specified time.")
		bedrock_listCustomModelsCmd.Flags().String("creation-time-before", "", "Return custom models created before the specified time.")
		bedrock_listCustomModelsCmd.Flags().String("foundation-model-arn-equals", "", "Return custom models only if the foundation model Amazon Resource Name (ARN) matches this parameter.")
		bedrock_listCustomModelsCmd.Flags().Bool("is-owned", false, "Return custom models depending on if the current account owns them (`true`) or if they were shared with the current account (`false`).")
		bedrock_listCustomModelsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrock_listCustomModelsCmd.Flags().String("model-status", "", "The status of them model to filter results by.")
		bedrock_listCustomModelsCmd.Flags().String("name-contains", "", "Return custom models only if the job name contains these characters.")
		bedrock_listCustomModelsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrock_listCustomModelsCmd.Flags().Bool("no-is-owned", false, "Return custom models depending on if the current account owns them (`true`) or if they were shared with the current account (`false`).")
		bedrock_listCustomModelsCmd.Flags().String("sort-by", "", "The field to sort by in the returned list of models.")
		bedrock_listCustomModelsCmd.Flags().String("sort-order", "", "The sort order of the results.")
		bedrock_listCustomModelsCmd.Flag("no-is-owned").Hidden = true
	})
	bedrockCmd.AddCommand(bedrock_listCustomModelsCmd)
}
