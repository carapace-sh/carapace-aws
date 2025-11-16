package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Changes the settings of an existing Cloud9 development environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_updateEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_updateEnvironmentCmd).Standalone()

		cloud9_updateEnvironmentCmd.Flags().String("description", "", "Any new or replacement description for the environment.")
		cloud9_updateEnvironmentCmd.Flags().String("environment-id", "", "The ID of the environment to change settings.")
		cloud9_updateEnvironmentCmd.Flags().String("managed-credentials-action", "", "Allows the environment owner to turn on or turn off the Amazon Web Services managed temporary credentials for an Cloud9 environment by using one of the following values:")
		cloud9_updateEnvironmentCmd.Flags().String("name", "", "A replacement name for the environment.")
		cloud9_updateEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	cloud9Cmd.AddCommand(cloud9_updateEnvironmentCmd)
}
