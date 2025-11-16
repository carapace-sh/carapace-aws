package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_deleteBlueprintCmd = &cobra.Command{
	Use:   "delete-blueprint",
	Short: "Deletes an existing Amazon Bedrock Data Automation Blueprint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_deleteBlueprintCmd).Standalone()

	bedrockDataAutomation_deleteBlueprintCmd.Flags().String("blueprint-arn", "", "ARN generated at the server side when a Blueprint is created")
	bedrockDataAutomation_deleteBlueprintCmd.Flags().String("blueprint-version", "", "Optional field to delete a specific Blueprint version")
	bedrockDataAutomation_deleteBlueprintCmd.MarkFlagRequired("blueprint-arn")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_deleteBlueprintCmd)
}
