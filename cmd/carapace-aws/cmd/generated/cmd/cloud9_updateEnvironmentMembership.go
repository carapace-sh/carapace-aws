package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_updateEnvironmentMembershipCmd = &cobra.Command{
	Use:   "update-environment-membership",
	Short: "Changes the settings of an existing environment member for an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_updateEnvironmentMembershipCmd).Standalone()

	cloud9_updateEnvironmentMembershipCmd.Flags().String("environment-id", "", "The ID of the environment for the environment member whose settings you want to change.")
	cloud9_updateEnvironmentMembershipCmd.Flags().String("permissions", "", "The replacement type of environment member permissions you want to associate with this environment member.")
	cloud9_updateEnvironmentMembershipCmd.Flags().String("user-arn", "", "The Amazon Resource Name (ARN) of the environment member whose settings you want to change.")
	cloud9_updateEnvironmentMembershipCmd.MarkFlagRequired("environment-id")
	cloud9_updateEnvironmentMembershipCmd.MarkFlagRequired("permissions")
	cloud9_updateEnvironmentMembershipCmd.MarkFlagRequired("user-arn")
	cloud9Cmd.AddCommand(cloud9_updateEnvironmentMembershipCmd)
}
