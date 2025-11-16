package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomationRuntimeCmd = &cobra.Command{
	Use:   "bedrock-data-automation-runtime",
	Short: "Amazon Bedrock Data Automation Runtime",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomationRuntimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomationRuntimeCmd).Standalone()

	})
	rootCmd.AddCommand(bedrockDataAutomationRuntimeCmd)
}
