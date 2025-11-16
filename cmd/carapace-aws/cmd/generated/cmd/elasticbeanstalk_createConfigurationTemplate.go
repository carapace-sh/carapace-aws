package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_createConfigurationTemplateCmd = &cobra.Command{
	Use:   "create-configuration-template",
	Short: "Creates an AWS Elastic Beanstalk configuration template, associated with a specific Elastic Beanstalk application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_createConfigurationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_createConfigurationTemplateCmd).Standalone()

		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("application-name", "", "The name of the Elastic Beanstalk application to associate with this configuration template.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("description", "", "An optional description for this configuration.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("environment-id", "", "The ID of an environment whose settings you want to use to create the configuration template.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("option-settings", "", "Option values for the Elastic Beanstalk configuration, such as the instance type.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("platform-arn", "", "The Amazon Resource Name (ARN) of the custom platform.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("solution-stack-name", "", "The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("source-configuration", "", "An Elastic Beanstalk configuration template to base this one on.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("tags", "", "Specifies the tags applied to the configuration template.")
		elasticbeanstalk_createConfigurationTemplateCmd.Flags().String("template-name", "", "The name of the configuration template.")
		elasticbeanstalk_createConfigurationTemplateCmd.MarkFlagRequired("application-name")
		elasticbeanstalk_createConfigurationTemplateCmd.MarkFlagRequired("template-name")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_createConfigurationTemplateCmd)
}
