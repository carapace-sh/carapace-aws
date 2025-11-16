package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_createDataAutomationProjectCmd = &cobra.Command{
	Use:   "create-data-automation-project",
	Short: "Creates an Amazon Bedrock Data Automation Project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_createDataAutomationProjectCmd).Standalone()

	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("client-token", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("custom-output-configuration", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("encryption-configuration", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("override-configuration", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("project-description", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("project-name", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("project-stage", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("standard-output-configuration", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.Flags().String("tags", "", "")
	bedrockDataAutomation_createDataAutomationProjectCmd.MarkFlagRequired("project-name")
	bedrockDataAutomation_createDataAutomationProjectCmd.MarkFlagRequired("standard-output-configuration")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_createDataAutomationProjectCmd)
}
