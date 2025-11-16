package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_updateDataAutomationProjectCmd = &cobra.Command{
	Use:   "update-data-automation-project",
	Short: "Updates an existing Amazon Bedrock Data Automation Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_updateDataAutomationProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomation_updateDataAutomationProjectCmd).Standalone()

		bedrockDataAutomation_updateDataAutomationProjectCmd.Flags().String("custom-output-configuration", "", "")
		bedrockDataAutomation_updateDataAutomationProjectCmd.Flags().String("encryption-configuration", "", "")
		bedrockDataAutomation_updateDataAutomationProjectCmd.Flags().String("override-configuration", "", "")
		bedrockDataAutomation_updateDataAutomationProjectCmd.Flags().String("project-arn", "", "ARN generated at the server side when a DataAutomationProject is created")
		bedrockDataAutomation_updateDataAutomationProjectCmd.Flags().String("project-description", "", "")
		bedrockDataAutomation_updateDataAutomationProjectCmd.Flags().String("project-stage", "", "")
		bedrockDataAutomation_updateDataAutomationProjectCmd.Flags().String("standard-output-configuration", "", "")
		bedrockDataAutomation_updateDataAutomationProjectCmd.MarkFlagRequired("project-arn")
		bedrockDataAutomation_updateDataAutomationProjectCmd.MarkFlagRequired("standard-output-configuration")
	})
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_updateDataAutomationProjectCmd)
}
