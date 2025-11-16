package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_describeEnvironmentStatusCmd = &cobra.Command{
	Use:   "describe-environment-status",
	Short: "Gets status information for an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_describeEnvironmentStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_describeEnvironmentStatusCmd).Standalone()

		cloud9_describeEnvironmentStatusCmd.Flags().String("environment-id", "", "The ID of the environment to get status information about.")
		cloud9_describeEnvironmentStatusCmd.MarkFlagRequired("environment-id")
	})
	cloud9Cmd.AddCommand(cloud9_describeEnvironmentStatusCmd)
}
