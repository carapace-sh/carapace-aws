package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteEnvironmentActionCmd = &cobra.Command{
	Use:   "delete-environment-action",
	Short: "Deletes an action for the environment, for example, deletes a console link for an analytics tool that is available in this environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteEnvironmentActionCmd).Standalone()

	datazone_deleteEnvironmentActionCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which an environment action is deleted.")
	datazone_deleteEnvironmentActionCmd.Flags().String("environment-identifier", "", "The ID of the environment where an environment action is deleted.")
	datazone_deleteEnvironmentActionCmd.Flags().String("identifier", "", "The ID of the environment action that is deleted.")
	datazone_deleteEnvironmentActionCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteEnvironmentActionCmd.MarkFlagRequired("environment-identifier")
	datazone_deleteEnvironmentActionCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteEnvironmentActionCmd)
}
