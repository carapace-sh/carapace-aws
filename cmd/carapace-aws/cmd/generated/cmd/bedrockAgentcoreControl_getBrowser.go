package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getBrowserCmd = &cobra.Command{
	Use:   "get-browser",
	Short: "Gets information about a custom browser.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getBrowserCmd).Standalone()

	bedrockAgentcoreControl_getBrowserCmd.Flags().String("browser-id", "", "The unique identifier of the browser to retrieve.")
	bedrockAgentcoreControl_getBrowserCmd.MarkFlagRequired("browser-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getBrowserCmd)
}
