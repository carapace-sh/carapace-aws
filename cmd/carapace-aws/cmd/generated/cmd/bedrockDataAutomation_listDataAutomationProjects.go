package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_listDataAutomationProjectsCmd = &cobra.Command{
	Use:   "list-data-automation-projects",
	Short: "Lists all existing Amazon Bedrock Data Automation Projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_listDataAutomationProjectsCmd).Standalone()

	bedrockDataAutomation_listDataAutomationProjectsCmd.Flags().String("blueprint-filter", "", "")
	bedrockDataAutomation_listDataAutomationProjectsCmd.Flags().String("max-results", "", "")
	bedrockDataAutomation_listDataAutomationProjectsCmd.Flags().String("next-token", "", "")
	bedrockDataAutomation_listDataAutomationProjectsCmd.Flags().String("project-stage-filter", "", "")
	bedrockDataAutomation_listDataAutomationProjectsCmd.Flags().String("resource-owner", "", "")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_listDataAutomationProjectsCmd)
}
