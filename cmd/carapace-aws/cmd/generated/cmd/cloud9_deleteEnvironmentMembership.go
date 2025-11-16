package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_deleteEnvironmentMembershipCmd = &cobra.Command{
	Use:   "delete-environment-membership",
	Short: "Deletes an environment member from a development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_deleteEnvironmentMembershipCmd).Standalone()

	cloud9_deleteEnvironmentMembershipCmd.Flags().String("environment-id", "", "The ID of the environment to delete the environment member from.")
	cloud9_deleteEnvironmentMembershipCmd.Flags().String("user-arn", "", "The Amazon Resource Name (ARN) of the environment member to delete from the environment.")
	cloud9_deleteEnvironmentMembershipCmd.MarkFlagRequired("environment-id")
	cloud9_deleteEnvironmentMembershipCmd.MarkFlagRequired("user-arn")
	cloud9Cmd.AddCommand(cloud9_deleteEnvironmentMembershipCmd)
}
