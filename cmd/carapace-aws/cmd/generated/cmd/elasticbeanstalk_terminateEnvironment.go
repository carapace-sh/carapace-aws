package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_terminateEnvironmentCmd = &cobra.Command{
	Use:   "terminate-environment",
	Short: "Terminates the specified environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_terminateEnvironmentCmd).Standalone()

	elasticbeanstalk_terminateEnvironmentCmd.Flags().String("environment-id", "", "The ID of the environment to terminate.")
	elasticbeanstalk_terminateEnvironmentCmd.Flags().String("environment-name", "", "The name of the environment to terminate.")
	elasticbeanstalk_terminateEnvironmentCmd.Flags().String("force-terminate", "", "Terminates the target environment even if another environment in the same group is dependent on it.")
	elasticbeanstalk_terminateEnvironmentCmd.Flags().String("terminate-resources", "", "Indicates whether the associated AWS resources should shut down when the environment is terminated:")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_terminateEnvironmentCmd)
}
