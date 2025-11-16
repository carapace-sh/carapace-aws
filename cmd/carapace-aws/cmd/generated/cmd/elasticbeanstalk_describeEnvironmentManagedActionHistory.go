package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd = &cobra.Command{
	Use:   "describe-environment-managed-action-history",
	Short: "Lists an environment's completed and failed managed actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd).Standalone()

		elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd.Flags().String("environment-id", "", "The environment ID of the target environment.")
		elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd.Flags().String("environment-name", "", "The name of the target environment.")
		elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd.Flags().String("max-items", "", "The maximum number of items to return for a single request.")
		elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd.Flags().String("next-token", "", "The pagination token returned by a previous request.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeEnvironmentManagedActionHistoryCmd)
}
