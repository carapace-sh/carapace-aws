package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_listBlueprintsCmd = &cobra.Command{
	Use:   "list-blueprints",
	Short: "Lists all existing Amazon Bedrock Data Automation Blueprints",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_listBlueprintsCmd).Standalone()

	bedrockDataAutomation_listBlueprintsCmd.Flags().String("blueprint-arn", "", "")
	bedrockDataAutomation_listBlueprintsCmd.Flags().String("blueprint-stage-filter", "", "")
	bedrockDataAutomation_listBlueprintsCmd.Flags().String("max-results", "", "")
	bedrockDataAutomation_listBlueprintsCmd.Flags().String("next-token", "", "")
	bedrockDataAutomation_listBlueprintsCmd.Flags().String("project-filter", "", "")
	bedrockDataAutomation_listBlueprintsCmd.Flags().String("resource-owner", "", "")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_listBlueprintsCmd)
}
