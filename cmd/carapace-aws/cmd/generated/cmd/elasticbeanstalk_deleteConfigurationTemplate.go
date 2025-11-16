package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_deleteConfigurationTemplateCmd = &cobra.Command{
	Use:   "delete-configuration-template",
	Short: "Deletes the specified configuration template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_deleteConfigurationTemplateCmd).Standalone()

	elasticbeanstalk_deleteConfigurationTemplateCmd.Flags().String("application-name", "", "The name of the application to delete the configuration template from.")
	elasticbeanstalk_deleteConfigurationTemplateCmd.Flags().String("template-name", "", "The name of the configuration template to delete.")
	elasticbeanstalk_deleteConfigurationTemplateCmd.MarkFlagRequired("application-name")
	elasticbeanstalk_deleteConfigurationTemplateCmd.MarkFlagRequired("template-name")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_deleteConfigurationTemplateCmd)
}
