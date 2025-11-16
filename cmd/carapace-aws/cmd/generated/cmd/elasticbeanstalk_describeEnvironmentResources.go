package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeEnvironmentResourcesCmd = &cobra.Command{
	Use:   "describe-environment-resources",
	Short: "Returns AWS resources for this environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeEnvironmentResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeEnvironmentResourcesCmd).Standalone()

		elasticbeanstalk_describeEnvironmentResourcesCmd.Flags().String("environment-id", "", "The ID of the environment to retrieve AWS resource usage data.")
		elasticbeanstalk_describeEnvironmentResourcesCmd.Flags().String("environment-name", "", "The name of the environment to retrieve AWS resource usage data.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeEnvironmentResourcesCmd)
}
