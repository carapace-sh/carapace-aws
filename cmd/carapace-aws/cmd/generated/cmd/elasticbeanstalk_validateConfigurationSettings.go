package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_validateConfigurationSettingsCmd = &cobra.Command{
	Use:   "validate-configuration-settings",
	Short: "Takes a set of configuration settings and either a configuration template or environment, and determines whether those values are valid.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_validateConfigurationSettingsCmd).Standalone()

	elasticbeanstalk_validateConfigurationSettingsCmd.Flags().String("application-name", "", "The name of the application that the configuration template or environment belongs to.")
	elasticbeanstalk_validateConfigurationSettingsCmd.Flags().String("environment-name", "", "The name of the environment to validate the settings against.")
	elasticbeanstalk_validateConfigurationSettingsCmd.Flags().String("option-settings", "", "A list of the options and desired values to evaluate.")
	elasticbeanstalk_validateConfigurationSettingsCmd.Flags().String("template-name", "", "The name of the configuration template to validate the settings against.")
	elasticbeanstalk_validateConfigurationSettingsCmd.MarkFlagRequired("application-name")
	elasticbeanstalk_validateConfigurationSettingsCmd.MarkFlagRequired("option-settings")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_validateConfigurationSettingsCmd)
}
