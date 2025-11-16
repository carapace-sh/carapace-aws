package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getEnvironmentActionCmd = &cobra.Command{
	Use:   "get-environment-action",
	Short: "Gets the specified environment action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getEnvironmentActionCmd).Standalone()

	datazone_getEnvironmentActionCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the `GetEnvironmentAction` API is invoked.")
	datazone_getEnvironmentActionCmd.Flags().String("environment-identifier", "", "The environment ID of the environment action.")
	datazone_getEnvironmentActionCmd.Flags().String("identifier", "", "The ID of the environment action")
	datazone_getEnvironmentActionCmd.MarkFlagRequired("domain-identifier")
	datazone_getEnvironmentActionCmd.MarkFlagRequired("environment-identifier")
	datazone_getEnvironmentActionCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getEnvironmentActionCmd)
}
