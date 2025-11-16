package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listInferenceProfilesCmd = &cobra.Command{
	Use:   "list-inference-profiles",
	Short: "Returns a list of inference profiles that you can use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listInferenceProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_listInferenceProfilesCmd).Standalone()

		bedrock_listInferenceProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrock_listInferenceProfilesCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrock_listInferenceProfilesCmd.Flags().String("type-equals", "", "Filters for inference profiles that match the type you specify.")
	})
	bedrockCmd.AddCommand(bedrock_listInferenceProfilesCmd)
}
