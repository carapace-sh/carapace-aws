package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createEnvironmentActionCmd = &cobra.Command{
	Use:   "create-environment-action",
	Short: "Creates an action for the environment, for example, creates a console link for an analytics tool that is available in this environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createEnvironmentActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createEnvironmentActionCmd).Standalone()

		datazone_createEnvironmentActionCmd.Flags().String("description", "", "The description of the environment action that is being created in the environment.")
		datazone_createEnvironmentActionCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the environment action is created.")
		datazone_createEnvironmentActionCmd.Flags().String("environment-identifier", "", "The ID of the environment in which the environment action is created.")
		datazone_createEnvironmentActionCmd.Flags().String("name", "", "The name of the environment action.")
		datazone_createEnvironmentActionCmd.Flags().String("parameters", "", "The parameters of the environment action.")
		datazone_createEnvironmentActionCmd.MarkFlagRequired("domain-identifier")
		datazone_createEnvironmentActionCmd.MarkFlagRequired("environment-identifier")
		datazone_createEnvironmentActionCmd.MarkFlagRequired("name")
		datazone_createEnvironmentActionCmd.MarkFlagRequired("parameters")
	})
	datazoneCmd.AddCommand(datazone_createEnvironmentActionCmd)
}
