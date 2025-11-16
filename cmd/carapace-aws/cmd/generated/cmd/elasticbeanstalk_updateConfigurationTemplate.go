package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_updateConfigurationTemplateCmd = &cobra.Command{
	Use:   "update-configuration-template",
	Short: "Updates the specified configuration template to have the specified properties or configuration option values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_updateConfigurationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_updateConfigurationTemplateCmd).Standalone()

		elasticbeanstalk_updateConfigurationTemplateCmd.Flags().String("application-name", "", "The name of the application associated with the configuration template to update.")
		elasticbeanstalk_updateConfigurationTemplateCmd.Flags().String("description", "", "A new description for the configuration.")
		elasticbeanstalk_updateConfigurationTemplateCmd.Flags().String("option-settings", "", "A list of configuration option settings to update with the new specified option value.")
		elasticbeanstalk_updateConfigurationTemplateCmd.Flags().String("options-to-remove", "", "A list of configuration options to remove from the configuration set.")
		elasticbeanstalk_updateConfigurationTemplateCmd.Flags().String("template-name", "", "The name of the configuration template to update.")
		elasticbeanstalk_updateConfigurationTemplateCmd.MarkFlagRequired("application-name")
		elasticbeanstalk_updateConfigurationTemplateCmd.MarkFlagRequired("template-name")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_updateConfigurationTemplateCmd)
}
