package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeEnvironmentManagedActionsCmd = &cobra.Command{
	Use:   "describe-environment-managed-actions",
	Short: "Lists an environment's upcoming and in-progress managed actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeEnvironmentManagedActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeEnvironmentManagedActionsCmd).Standalone()

		elasticbeanstalk_describeEnvironmentManagedActionsCmd.Flags().String("environment-id", "", "The environment ID of the target environment.")
		elasticbeanstalk_describeEnvironmentManagedActionsCmd.Flags().String("environment-name", "", "The name of the target environment.")
		elasticbeanstalk_describeEnvironmentManagedActionsCmd.Flags().String("status", "", "To show only actions with a particular status, specify a status.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeEnvironmentManagedActionsCmd)
}
