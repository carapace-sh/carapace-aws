package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_getBlueprintCmd = &cobra.Command{
	Use:   "get-blueprint",
	Short: "Gets an existing Amazon Bedrock Data Automation Blueprint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_getBlueprintCmd).Standalone()

	bedrockDataAutomation_getBlueprintCmd.Flags().String("blueprint-arn", "", "ARN generated at the server side when a Blueprint is created")
	bedrockDataAutomation_getBlueprintCmd.Flags().String("blueprint-stage", "", "Optional field to get a specific Blueprint stage")
	bedrockDataAutomation_getBlueprintCmd.Flags().String("blueprint-version", "", "Optional field to get a specific Blueprint version")
	bedrockDataAutomation_getBlueprintCmd.MarkFlagRequired("blueprint-arn")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_getBlueprintCmd)
}
