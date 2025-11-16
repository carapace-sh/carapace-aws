package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_updateBrowserStreamCmd = &cobra.Command{
	Use:   "update-browser-stream",
	Short: "Updates a browser stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_updateBrowserStreamCmd).Standalone()

	bedrockAgentcore_updateBrowserStreamCmd.Flags().String("browser-identifier", "", "The identifier of the browser.")
	bedrockAgentcore_updateBrowserStreamCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
	bedrockAgentcore_updateBrowserStreamCmd.Flags().String("session-id", "", "The identifier of the browser session.")
	bedrockAgentcore_updateBrowserStreamCmd.Flags().String("stream-update", "", "The update to apply to the browser stream.")
	bedrockAgentcore_updateBrowserStreamCmd.MarkFlagRequired("browser-identifier")
	bedrockAgentcore_updateBrowserStreamCmd.MarkFlagRequired("session-id")
	bedrockAgentcore_updateBrowserStreamCmd.MarkFlagRequired("stream-update")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_updateBrowserStreamCmd)
}
