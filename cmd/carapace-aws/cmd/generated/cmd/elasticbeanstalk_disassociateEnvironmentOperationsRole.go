package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_disassociateEnvironmentOperationsRoleCmd = &cobra.Command{
	Use:   "disassociate-environment-operations-role",
	Short: "Disassociate the operations role from an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_disassociateEnvironmentOperationsRoleCmd).Standalone()

	elasticbeanstalk_disassociateEnvironmentOperationsRoleCmd.Flags().String("environment-name", "", "The name of the environment from which to disassociate the operations role.")
	elasticbeanstalk_disassociateEnvironmentOperationsRoleCmd.MarkFlagRequired("environment-name")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_disassociateEnvironmentOperationsRoleCmd)
}
