package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Updates the specified environment in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateEnvironmentCmd).Standalone()

		datazone_updateEnvironmentCmd.Flags().String("blueprint-version", "", "The blueprint version to which the environment should be updated.")
		datazone_updateEnvironmentCmd.Flags().String("description", "", "The description to be updated as part of the `UpdateEnvironment` action.")
		datazone_updateEnvironmentCmd.Flags().String("domain-identifier", "", "The identifier of the domain in which the environment is to be updated.")
		datazone_updateEnvironmentCmd.Flags().String("glossary-terms", "", "The glossary terms to be updated as part of the `UpdateEnvironment` action.")
		datazone_updateEnvironmentCmd.Flags().String("identifier", "", "The identifier of the environment that is to be updated.")
		datazone_updateEnvironmentCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateEnvironment` action.")
		datazone_updateEnvironmentCmd.Flags().String("user-parameters", "", "The user parameters of the environment.")
		datazone_updateEnvironmentCmd.MarkFlagRequired("domain-identifier")
		datazone_updateEnvironmentCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateEnvironmentCmd)
}
