package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listPromptRoutersCmd = &cobra.Command{
	Use:   "list-prompt-routers",
	Short: "Retrieves a list of prompt routers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listPromptRoutersCmd).Standalone()

	bedrock_listPromptRoutersCmd.Flags().String("max-results", "", "The maximum number of prompt routers to return in one page of results.")
	bedrock_listPromptRoutersCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	bedrock_listPromptRoutersCmd.Flags().String("type", "", "The type of the prompt routers, such as whether it's default or custom.")
	bedrockCmd.AddCommand(bedrock_listPromptRoutersCmd)
}
