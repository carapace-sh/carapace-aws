package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeConfigurationSettingsCmd = &cobra.Command{
	Use:   "describe-configuration-settings",
	Short: "Returns a description of the settings for the specified configuration set, that is, either a configuration template or the configuration set associated with a running environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeConfigurationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeConfigurationSettingsCmd).Standalone()

		elasticbeanstalk_describeConfigurationSettingsCmd.Flags().String("application-name", "", "The application for the environment or configuration template.")
		elasticbeanstalk_describeConfigurationSettingsCmd.Flags().String("environment-name", "", "The name of the environment to describe.")
		elasticbeanstalk_describeConfigurationSettingsCmd.Flags().String("template-name", "", "The name of the configuration template to describe.")
		elasticbeanstalk_describeConfigurationSettingsCmd.MarkFlagRequired("application-name")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeConfigurationSettingsCmd)
}
