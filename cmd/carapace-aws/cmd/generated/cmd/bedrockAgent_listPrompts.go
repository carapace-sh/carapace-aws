package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listPromptsCmd = &cobra.Command{
	Use:   "list-prompts",
	Short: "Returns either information about the working draft (`DRAFT` version) of each prompt in an account, or information about of all versions of a prompt, depending on whether you include the `promptIdentifier` field or not.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listPromptsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listPromptsCmd).Standalone()

		bedrockAgent_listPromptsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listPromptsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listPromptsCmd.Flags().String("prompt-identifier", "", "The unique identifier of the prompt for whose versions you want to return information.")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listPromptsCmd)
}
