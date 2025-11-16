package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createBrowserCmd = &cobra.Command{
	Use:   "create-browser",
	Short: "Creates a custom browser.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createBrowserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_createBrowserCmd).Standalone()

		bedrockAgentcoreControl_createBrowserCmd.Flags().String("browser-signing", "", "The browser signing configuration that enables cryptographic agent identification using HTTP message signatures for web bot authentication.")
		bedrockAgentcoreControl_createBrowserCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
		bedrockAgentcoreControl_createBrowserCmd.Flags().String("description", "", "The description of the browser.")
		bedrockAgentcoreControl_createBrowserCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the browser to access Amazon Web Services services.")
		bedrockAgentcoreControl_createBrowserCmd.Flags().String("name", "", "The name of the browser.")
		bedrockAgentcoreControl_createBrowserCmd.Flags().String("network-configuration", "", "The network configuration for the browser.")
		bedrockAgentcoreControl_createBrowserCmd.Flags().String("recording", "", "The recording configuration for the browser.")
		bedrockAgentcoreControl_createBrowserCmd.Flags().String("tags", "", "A map of tag keys and values to assign to the browser.")
		bedrockAgentcoreControl_createBrowserCmd.MarkFlagRequired("name")
		bedrockAgentcoreControl_createBrowserCmd.MarkFlagRequired("network-configuration")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createBrowserCmd)
}
