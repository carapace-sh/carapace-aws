package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes an Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_deleteEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_deleteEnvironmentCmd).Standalone()

		cloud9_deleteEnvironmentCmd.Flags().String("environment-id", "", "The ID of the environment to delete.")
		cloud9_deleteEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	cloud9Cmd.AddCommand(cloud9_deleteEnvironmentCmd)
}
