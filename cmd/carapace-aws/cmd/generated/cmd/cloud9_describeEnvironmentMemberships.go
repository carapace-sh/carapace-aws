package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_describeEnvironmentMembershipsCmd = &cobra.Command{
	Use:   "describe-environment-memberships",
	Short: "Gets information about environment members for an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_describeEnvironmentMembershipsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_describeEnvironmentMembershipsCmd).Standalone()

		cloud9_describeEnvironmentMembershipsCmd.Flags().String("environment-id", "", "The ID of the environment to get environment member information about.")
		cloud9_describeEnvironmentMembershipsCmd.Flags().String("max-results", "", "The maximum number of environment members to get information about.")
		cloud9_describeEnvironmentMembershipsCmd.Flags().String("next-token", "", "During a previous call, if there are more than 25 items in the list, only the first 25 items are returned, along with a unique string called a *next token*.")
		cloud9_describeEnvironmentMembershipsCmd.Flags().String("permissions", "", "The type of environment member permissions to get information about.")
		cloud9_describeEnvironmentMembershipsCmd.Flags().String("user-arn", "", "The Amazon Resource Name (ARN) of an individual environment member to get information about.")
	})
	cloud9Cmd.AddCommand(cloud9_describeEnvironmentMembershipsCmd)
}
