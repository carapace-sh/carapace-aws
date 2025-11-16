package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_createBlueprintVersionCmd = &cobra.Command{
	Use:   "create-blueprint-version",
	Short: "Creates a new version of an existing Amazon Bedrock Data Automation Blueprint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_createBlueprintVersionCmd).Standalone()

	bedrockDataAutomation_createBlueprintVersionCmd.Flags().String("blueprint-arn", "", "ARN generated at the server side when a Blueprint is created")
	bedrockDataAutomation_createBlueprintVersionCmd.Flags().String("client-token", "", "")
	bedrockDataAutomation_createBlueprintVersionCmd.MarkFlagRequired("blueprint-arn")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_createBlueprintVersionCmd)
}
