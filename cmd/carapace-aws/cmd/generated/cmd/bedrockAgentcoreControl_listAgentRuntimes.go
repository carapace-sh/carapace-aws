package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listAgentRuntimesCmd = &cobra.Command{
	Use:   "list-agent-runtimes",
	Short: "Lists all Amazon Secure Agents in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listAgentRuntimesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_listAgentRuntimesCmd).Standalone()

		bedrockAgentcoreControl_listAgentRuntimesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgentcoreControl_listAgentRuntimesCmd.Flags().String("next-token", "", "A token to retrieve the next page of results.")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listAgentRuntimesCmd)
}
