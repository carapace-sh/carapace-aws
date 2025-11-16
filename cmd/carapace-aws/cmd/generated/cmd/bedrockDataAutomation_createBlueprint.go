package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_createBlueprintCmd = &cobra.Command{
	Use:   "create-blueprint",
	Short: "Creates an Amazon Bedrock Data Automation Blueprint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_createBlueprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomation_createBlueprintCmd).Standalone()

		bedrockDataAutomation_createBlueprintCmd.Flags().String("blueprint-name", "", "")
		bedrockDataAutomation_createBlueprintCmd.Flags().String("blueprint-stage", "", "")
		bedrockDataAutomation_createBlueprintCmd.Flags().String("client-token", "", "")
		bedrockDataAutomation_createBlueprintCmd.Flags().String("encryption-configuration", "", "")
		bedrockDataAutomation_createBlueprintCmd.Flags().String("schema", "", "")
		bedrockDataAutomation_createBlueprintCmd.Flags().String("tags", "", "")
		bedrockDataAutomation_createBlueprintCmd.Flags().String("type", "", "")
		bedrockDataAutomation_createBlueprintCmd.MarkFlagRequired("blueprint-name")
		bedrockDataAutomation_createBlueprintCmd.MarkFlagRequired("schema")
		bedrockDataAutomation_createBlueprintCmd.MarkFlagRequired("type")
	})
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_createBlueprintCmd)
}
