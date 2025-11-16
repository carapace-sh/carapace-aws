package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_associateEnvironmentOperationsRoleCmd = &cobra.Command{
	Use:   "associate-environment-operations-role",
	Short: "Add or change the operations role used by an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_associateEnvironmentOperationsRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_associateEnvironmentOperationsRoleCmd).Standalone()

		elasticbeanstalk_associateEnvironmentOperationsRoleCmd.Flags().String("environment-name", "", "The name of the environment to which to set the operations role.")
		elasticbeanstalk_associateEnvironmentOperationsRoleCmd.Flags().String("operations-role", "", "The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.")
		elasticbeanstalk_associateEnvironmentOperationsRoleCmd.MarkFlagRequired("environment-name")
		elasticbeanstalk_associateEnvironmentOperationsRoleCmd.MarkFlagRequired("operations-role")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_associateEnvironmentOperationsRoleCmd)
}
