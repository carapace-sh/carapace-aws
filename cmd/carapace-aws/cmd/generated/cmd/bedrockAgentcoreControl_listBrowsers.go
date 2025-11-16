package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listBrowsersCmd = &cobra.Command{
	Use:   "list-browsers",
	Short: "Lists all custom browsers in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listBrowsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_listBrowsersCmd).Standalone()

		bedrockAgentcoreControl_listBrowsersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		bedrockAgentcoreControl_listBrowsersCmd.Flags().String("next-token", "", "The token for the next set of results.")
		bedrockAgentcoreControl_listBrowsersCmd.Flags().String("type", "", "The type of browsers to list.")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listBrowsersCmd)
}
