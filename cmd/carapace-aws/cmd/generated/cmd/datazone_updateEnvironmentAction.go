package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateEnvironmentActionCmd = &cobra.Command{
	Use:   "update-environment-action",
	Short: "Updates an environment action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateEnvironmentActionCmd).Standalone()

	datazone_updateEnvironmentActionCmd.Flags().String("description", "", "The description of the environment action.")
	datazone_updateEnvironmentActionCmd.Flags().String("domain-identifier", "", "The domain ID of the environment action.")
	datazone_updateEnvironmentActionCmd.Flags().String("environment-identifier", "", "The environment ID of the environment action.")
	datazone_updateEnvironmentActionCmd.Flags().String("identifier", "", "The ID of the environment action.")
	datazone_updateEnvironmentActionCmd.Flags().String("name", "", "The name of the environment action.")
	datazone_updateEnvironmentActionCmd.Flags().String("parameters", "", "The parameters of the environment action.")
	datazone_updateEnvironmentActionCmd.MarkFlagRequired("domain-identifier")
	datazone_updateEnvironmentActionCmd.MarkFlagRequired("environment-identifier")
	datazone_updateEnvironmentActionCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_updateEnvironmentActionCmd)
}
