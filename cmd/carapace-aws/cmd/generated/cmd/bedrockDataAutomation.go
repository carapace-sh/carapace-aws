package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomationCmd = &cobra.Command{
	Use:   "bedrock-data-automation",
	Short: "Amazon Bedrock Data Automation BuildTime",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomationCmd).Standalone()

	})
	rootCmd.AddCommand(bedrockDataAutomationCmd)
}
