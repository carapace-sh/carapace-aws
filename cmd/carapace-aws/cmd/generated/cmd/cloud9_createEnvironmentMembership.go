package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_createEnvironmentMembershipCmd = &cobra.Command{
	Use:   "create-environment-membership",
	Short: "Adds an environment member to an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_createEnvironmentMembershipCmd).Standalone()

	cloud9_createEnvironmentMembershipCmd.Flags().String("environment-id", "", "The ID of the environment that contains the environment member you want to add.")
	cloud9_createEnvironmentMembershipCmd.Flags().String("permissions", "", "The type of environment member permissions you want to associate with this environment member.")
	cloud9_createEnvironmentMembershipCmd.Flags().String("user-arn", "", "The Amazon Resource Name (ARN) of the environment member you want to add.")
	cloud9_createEnvironmentMembershipCmd.MarkFlagRequired("environment-id")
	cloud9_createEnvironmentMembershipCmd.MarkFlagRequired("permissions")
	cloud9_createEnvironmentMembershipCmd.MarkFlagRequired("user-arn")
	cloud9Cmd.AddCommand(cloud9_createEnvironmentMembershipCmd)
}
