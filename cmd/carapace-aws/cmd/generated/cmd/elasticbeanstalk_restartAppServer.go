package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_restartAppServerCmd = &cobra.Command{
	Use:   "restart-app-server",
	Short: "Causes the environment to restart the application container server running on each Amazon EC2 instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_restartAppServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_restartAppServerCmd).Standalone()

		elasticbeanstalk_restartAppServerCmd.Flags().String("environment-id", "", "The ID of the environment to restart the server for.")
		elasticbeanstalk_restartAppServerCmd.Flags().String("environment-name", "", "The name of the environment to restart the server for.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_restartAppServerCmd)
}
