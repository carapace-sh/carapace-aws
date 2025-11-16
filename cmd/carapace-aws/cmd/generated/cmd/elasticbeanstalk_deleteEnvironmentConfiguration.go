package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_deleteEnvironmentConfigurationCmd = &cobra.Command{
	Use:   "delete-environment-configuration",
	Short: "Deletes the draft configuration associated with the running environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_deleteEnvironmentConfigurationCmd).Standalone()

	elasticbeanstalk_deleteEnvironmentConfigurationCmd.Flags().String("application-name", "", "The name of the application the environment is associated with.")
	elasticbeanstalk_deleteEnvironmentConfigurationCmd.Flags().String("environment-name", "", "The name of the environment to delete the draft configuration from.")
	elasticbeanstalk_deleteEnvironmentConfigurationCmd.MarkFlagRequired("application-name")
	elasticbeanstalk_deleteEnvironmentConfigurationCmd.MarkFlagRequired("environment-name")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_deleteEnvironmentConfigurationCmd)
}
