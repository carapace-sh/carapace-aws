package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_updateBlueprintCmd = &cobra.Command{
	Use:   "update-blueprint",
	Short: "Updates an existing Amazon Bedrock Data Automation Blueprint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_updateBlueprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomation_updateBlueprintCmd).Standalone()

		bedrockDataAutomation_updateBlueprintCmd.Flags().String("blueprint-arn", "", "ARN generated at the server side when a Blueprint is created")
		bedrockDataAutomation_updateBlueprintCmd.Flags().String("blueprint-stage", "", "")
		bedrockDataAutomation_updateBlueprintCmd.Flags().String("encryption-configuration", "", "")
		bedrockDataAutomation_updateBlueprintCmd.Flags().String("schema", "", "")
		bedrockDataAutomation_updateBlueprintCmd.MarkFlagRequired("blueprint-arn")
		bedrockDataAutomation_updateBlueprintCmd.MarkFlagRequired("schema")
	})
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_updateBlueprintCmd)
}
