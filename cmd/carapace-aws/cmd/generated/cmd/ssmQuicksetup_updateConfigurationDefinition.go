package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_updateConfigurationDefinitionCmd = &cobra.Command{
	Use:   "update-configuration-definition",
	Short: "Updates a Quick Setup configuration definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_updateConfigurationDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_updateConfigurationDefinitionCmd).Standalone()

		ssmQuicksetup_updateConfigurationDefinitionCmd.Flags().String("id", "", "The ID of the configuration definition you want to update.")
		ssmQuicksetup_updateConfigurationDefinitionCmd.Flags().String("local-deployment-administration-role-arn", "", "The ARN of the IAM role used to administrate local configuration deployments.")
		ssmQuicksetup_updateConfigurationDefinitionCmd.Flags().String("local-deployment-execution-role-name", "", "The name of the IAM role used to deploy local configurations.")
		ssmQuicksetup_updateConfigurationDefinitionCmd.Flags().String("manager-arn", "", "The ARN of the configuration manager associated with the definition to update.")
		ssmQuicksetup_updateConfigurationDefinitionCmd.Flags().String("parameters", "", "The parameters for the configuration definition type.")
		ssmQuicksetup_updateConfigurationDefinitionCmd.Flags().String("type-version", "", "The version of the Quick Setup type to use.")
		ssmQuicksetup_updateConfigurationDefinitionCmd.MarkFlagRequired("id")
		ssmQuicksetup_updateConfigurationDefinitionCmd.MarkFlagRequired("manager-arn")
	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_updateConfigurationDefinitionCmd)
}
