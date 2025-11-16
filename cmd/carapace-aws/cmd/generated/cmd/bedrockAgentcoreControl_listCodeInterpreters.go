package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listCodeInterpretersCmd = &cobra.Command{
	Use:   "list-code-interpreters",
	Short: "Lists all custom code interpreters in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listCodeInterpretersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_listCodeInterpretersCmd).Standalone()

		bedrockAgentcoreControl_listCodeInterpretersCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgentcoreControl_listCodeInterpretersCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
		bedrockAgentcoreControl_listCodeInterpretersCmd.Flags().String("type", "", "The type of code interpreters to list.")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listCodeInterpretersCmd)
}
