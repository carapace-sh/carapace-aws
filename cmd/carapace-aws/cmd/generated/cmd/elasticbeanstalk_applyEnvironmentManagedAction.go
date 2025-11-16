package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_applyEnvironmentManagedActionCmd = &cobra.Command{
	Use:   "apply-environment-managed-action",
	Short: "Applies a scheduled managed action immediately.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_applyEnvironmentManagedActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_applyEnvironmentManagedActionCmd).Standalone()

		elasticbeanstalk_applyEnvironmentManagedActionCmd.Flags().String("action-id", "", "The action ID of the scheduled managed action to execute.")
		elasticbeanstalk_applyEnvironmentManagedActionCmd.Flags().String("environment-id", "", "The environment ID of the target environment.")
		elasticbeanstalk_applyEnvironmentManagedActionCmd.Flags().String("environment-name", "", "The name of the target environment.")
		elasticbeanstalk_applyEnvironmentManagedActionCmd.MarkFlagRequired("action-id")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_applyEnvironmentManagedActionCmd)
}
