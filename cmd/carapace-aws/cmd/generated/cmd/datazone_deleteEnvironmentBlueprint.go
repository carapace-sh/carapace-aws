package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteEnvironmentBlueprintCmd = &cobra.Command{
	Use:   "delete-environment-blueprint",
	Short: "Deletes a blueprint in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteEnvironmentBlueprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteEnvironmentBlueprintCmd).Standalone()

		datazone_deleteEnvironmentBlueprintCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the blueprint is deleted.")
		datazone_deleteEnvironmentBlueprintCmd.Flags().String("identifier", "", "The ID of the blueprint that is deleted.")
		datazone_deleteEnvironmentBlueprintCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteEnvironmentBlueprintCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteEnvironmentBlueprintCmd)
}
