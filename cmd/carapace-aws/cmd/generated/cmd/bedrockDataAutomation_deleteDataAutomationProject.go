package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_deleteDataAutomationProjectCmd = &cobra.Command{
	Use:   "delete-data-automation-project",
	Short: "Deletes an existing Amazon Bedrock Data Automation Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_deleteDataAutomationProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomation_deleteDataAutomationProjectCmd).Standalone()

		bedrockDataAutomation_deleteDataAutomationProjectCmd.Flags().String("project-arn", "", "ARN generated at the server side when a DataAutomationProject is created")
		bedrockDataAutomation_deleteDataAutomationProjectCmd.MarkFlagRequired("project-arn")
	})
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_deleteDataAutomationProjectCmd)
}
