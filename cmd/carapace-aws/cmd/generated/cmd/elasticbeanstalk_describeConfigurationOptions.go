package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeConfigurationOptionsCmd = &cobra.Command{
	Use:   "describe-configuration-options",
	Short: "Describes the configuration options that are used in a particular configuration template or environment, or that a specified solution stack defines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeConfigurationOptionsCmd).Standalone()

	elasticbeanstalk_describeConfigurationOptionsCmd.Flags().String("application-name", "", "The name of the application associated with the configuration template or environment.")
	elasticbeanstalk_describeConfigurationOptionsCmd.Flags().String("environment-name", "", "The name of the environment whose configuration options you want to describe.")
	elasticbeanstalk_describeConfigurationOptionsCmd.Flags().String("options", "", "If specified, restricts the descriptions to only the specified options.")
	elasticbeanstalk_describeConfigurationOptionsCmd.Flags().String("platform-arn", "", "The ARN of the custom platform.")
	elasticbeanstalk_describeConfigurationOptionsCmd.Flags().String("solution-stack-name", "", "The name of the solution stack whose configuration options you want to describe.")
	elasticbeanstalk_describeConfigurationOptionsCmd.Flags().String("template-name", "", "The name of the configuration template whose configuration options you want to describe.")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeConfigurationOptionsCmd)
}
