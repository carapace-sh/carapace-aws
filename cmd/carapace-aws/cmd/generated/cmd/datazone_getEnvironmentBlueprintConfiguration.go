package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getEnvironmentBlueprintConfigurationCmd = &cobra.Command{
	Use:   "get-environment-blueprint-configuration",
	Short: "Gets the blueprint configuration in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getEnvironmentBlueprintConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getEnvironmentBlueprintConfigurationCmd).Standalone()

		datazone_getEnvironmentBlueprintConfigurationCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where this blueprint exists.")
		datazone_getEnvironmentBlueprintConfigurationCmd.Flags().String("environment-blueprint-identifier", "", "He ID of the blueprint.")
		datazone_getEnvironmentBlueprintConfigurationCmd.MarkFlagRequired("domain-identifier")
		datazone_getEnvironmentBlueprintConfigurationCmd.MarkFlagRequired("environment-blueprint-identifier")
	})
	datazoneCmd.AddCommand(datazone_getEnvironmentBlueprintConfigurationCmd)
}
