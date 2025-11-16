package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteBrowserCmd = &cobra.Command{
	Use:   "delete-browser",
	Short: "Deletes a custom browser.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteBrowserCmd).Standalone()

	bedrockAgentcoreControl_deleteBrowserCmd.Flags().String("browser-id", "", "The unique identifier of the browser to delete.")
	bedrockAgentcoreControl_deleteBrowserCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	bedrockAgentcoreControl_deleteBrowserCmd.MarkFlagRequired("browser-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteBrowserCmd)
}
