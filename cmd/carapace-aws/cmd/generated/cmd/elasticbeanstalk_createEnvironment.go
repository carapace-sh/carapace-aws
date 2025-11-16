package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Launches an AWS Elastic Beanstalk environment for the specified application using the specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_createEnvironmentCmd).Standalone()

	elasticbeanstalk_createEnvironmentCmd.Flags().String("application-name", "", "The name of the application that is associated with this environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("cnameprefix", "", "If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("description", "", "Your description for this environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("environment-name", "", "A unique name for the environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("group-name", "", "The name of the group to which the target environment belongs.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("operations-role", "", "The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("option-settings", "", "If specified, AWS Elastic Beanstalk sets the specified configuration options to the requested value in the configuration set for the new environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("options-to-remove", "", "A list of custom user-defined configuration options to remove from the configuration set for this new environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("platform-arn", "", "The Amazon Resource Name (ARN) of the custom platform to use with the environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("solution-stack-name", "", "The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("tags", "", "Specifies the tags applied to resources in the environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("template-name", "", "The name of the Elastic Beanstalk configuration template to use with the environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("tier", "", "Specifies the tier to use in creating this environment.")
	elasticbeanstalk_createEnvironmentCmd.Flags().String("version-label", "", "The name of the application version to deploy.")
	elasticbeanstalk_createEnvironmentCmd.MarkFlagRequired("application-name")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_createEnvironmentCmd)
}
