package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_rebuildEnvironmentCmd = &cobra.Command{
	Use:   "rebuild-environment",
	Short: "Deletes and recreates all of the AWS resources (for example: the Auto Scaling group, load balancer, etc.) for a specified environment and forces a restart.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_rebuildEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_rebuildEnvironmentCmd).Standalone()

		elasticbeanstalk_rebuildEnvironmentCmd.Flags().String("environment-id", "", "The ID of the environment to rebuild.")
		elasticbeanstalk_rebuildEnvironmentCmd.Flags().String("environment-name", "", "The name of the environment to rebuild.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_rebuildEnvironmentCmd)
}
