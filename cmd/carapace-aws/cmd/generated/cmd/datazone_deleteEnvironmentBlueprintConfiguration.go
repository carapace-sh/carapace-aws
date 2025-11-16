package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteEnvironmentBlueprintConfigurationCmd = &cobra.Command{
	Use:   "delete-environment-blueprint-configuration",
	Short: "Deletes the blueprint configuration in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteEnvironmentBlueprintConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteEnvironmentBlueprintConfigurationCmd).Standalone()

		datazone_deleteEnvironmentBlueprintConfigurationCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the blueprint configuration is deleted.")
		datazone_deleteEnvironmentBlueprintConfigurationCmd.Flags().String("environment-blueprint-identifier", "", "The ID of the blueprint the configuration of which is deleted.")
		datazone_deleteEnvironmentBlueprintConfigurationCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteEnvironmentBlueprintConfigurationCmd.MarkFlagRequired("environment-blueprint-identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteEnvironmentBlueprintConfigurationCmd)
}
