package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_abortEnvironmentUpdateCmd = &cobra.Command{
	Use:   "abort-environment-update",
	Short: "Cancels in-progress environment configuration update or application version deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_abortEnvironmentUpdateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_abortEnvironmentUpdateCmd).Standalone()

		elasticbeanstalk_abortEnvironmentUpdateCmd.Flags().String("environment-id", "", "This specifies the ID of the environment with the in-progress update that you want to cancel.")
		elasticbeanstalk_abortEnvironmentUpdateCmd.Flags().String("environment-name", "", "This specifies the name of the environment with the in-progress update that you want to cancel.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_abortEnvironmentUpdateCmd)
}
