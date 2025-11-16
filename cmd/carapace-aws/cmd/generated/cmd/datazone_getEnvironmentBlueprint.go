package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getEnvironmentBlueprintCmd = &cobra.Command{
	Use:   "get-environment-blueprint",
	Short: "Gets an Amazon DataZone blueprint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getEnvironmentBlueprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getEnvironmentBlueprintCmd).Standalone()

		datazone_getEnvironmentBlueprintCmd.Flags().String("domain-identifier", "", "The identifier of the domain in which this blueprint exists.")
		datazone_getEnvironmentBlueprintCmd.Flags().String("identifier", "", "The ID of this Amazon DataZone blueprint.")
		datazone_getEnvironmentBlueprintCmd.MarkFlagRequired("domain-identifier")
		datazone_getEnvironmentBlueprintCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getEnvironmentBlueprintCmd)
}
