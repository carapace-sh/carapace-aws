package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Updates the environment description, deploys a new application version, updates the configuration settings to an entirely new configuration template, or updates select configuration option values in the running environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_updateEnvironmentCmd).Standalone()

	elasticbeanstalk_updateEnvironmentCmd.Flags().String("application-name", "", "The name of the application with which the environment is associated.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("description", "", "If this parameter is specified, AWS Elastic Beanstalk updates the description of this environment.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("environment-id", "", "The ID of the environment to update.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("environment-name", "", "The name of the environment to update.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("group-name", "", "The name of the group to which the target environment belongs.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("option-settings", "", "If specified, AWS Elastic Beanstalk updates the configuration set associated with the running environment and sets the specified configuration options to the requested value.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("options-to-remove", "", "A list of custom user-defined configuration options to remove from the configuration set for this environment.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("platform-arn", "", "The ARN of the platform, if used.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("solution-stack-name", "", "This specifies the platform version that the environment will run after the environment is updated.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("template-name", "", "If this parameter is specified, AWS Elastic Beanstalk deploys this configuration template to the environment.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("tier", "", "This specifies the tier to use to update the environment.")
	elasticbeanstalk_updateEnvironmentCmd.Flags().String("version-label", "", "If this parameter is specified, AWS Elastic Beanstalk deploys the named application version to the environment.")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_updateEnvironmentCmd)
}
