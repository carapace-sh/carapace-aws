package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateEnvironmentBlueprintCmd = &cobra.Command{
	Use:   "update-environment-blueprint",
	Short: "Updates an environment blueprint in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateEnvironmentBlueprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateEnvironmentBlueprintCmd).Standalone()

		datazone_updateEnvironmentBlueprintCmd.Flags().String("description", "", "The description to be updated as part of the `UpdateEnvironmentBlueprint` action.")
		datazone_updateEnvironmentBlueprintCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which an environment blueprint is to be updated.")
		datazone_updateEnvironmentBlueprintCmd.Flags().String("identifier", "", "The identifier of the environment blueprint to be updated.")
		datazone_updateEnvironmentBlueprintCmd.Flags().String("provisioning-properties", "", "The provisioning properties to be updated as part of the `UpdateEnvironmentBlueprint` action.")
		datazone_updateEnvironmentBlueprintCmd.Flags().String("user-parameters", "", "The user parameters to be updated as part of the `UpdateEnvironmentBlueprint` action.")
		datazone_updateEnvironmentBlueprintCmd.MarkFlagRequired("domain-identifier")
		datazone_updateEnvironmentBlueprintCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateEnvironmentBlueprintCmd)
}
