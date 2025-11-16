package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_getDataAutomationProjectCmd = &cobra.Command{
	Use:   "get-data-automation-project",
	Short: "Gets an existing Amazon Bedrock Data Automation Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_getDataAutomationProjectCmd).Standalone()

	bedrockDataAutomation_getDataAutomationProjectCmd.Flags().String("project-arn", "", "ARN generated at the server side when a DataAutomationProject is created")
	bedrockDataAutomation_getDataAutomationProjectCmd.Flags().String("project-stage", "", "Optional field to delete a specific DataAutomationProject stage")
	bedrockDataAutomation_getDataAutomationProjectCmd.MarkFlagRequired("project-arn")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_getDataAutomationProjectCmd)
}
