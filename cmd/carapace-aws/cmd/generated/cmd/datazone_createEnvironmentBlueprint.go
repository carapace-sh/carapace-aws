package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createEnvironmentBlueprintCmd = &cobra.Command{
	Use:   "create-environment-blueprint",
	Short: "Creates a Amazon DataZone blueprint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createEnvironmentBlueprintCmd).Standalone()

	datazone_createEnvironmentBlueprintCmd.Flags().String("description", "", "The description of the Amazon DataZone blueprint.")
	datazone_createEnvironmentBlueprintCmd.Flags().String("domain-identifier", "", "The identifier of the domain in which this blueprint is created.")
	datazone_createEnvironmentBlueprintCmd.Flags().String("name", "", "The name of this Amazon DataZone blueprint.")
	datazone_createEnvironmentBlueprintCmd.Flags().String("provisioning-properties", "", "The provisioning properties of this Amazon DataZone blueprint.")
	datazone_createEnvironmentBlueprintCmd.Flags().String("user-parameters", "", "The user parameters of this Amazon DataZone blueprint.")
	datazone_createEnvironmentBlueprintCmd.MarkFlagRequired("domain-identifier")
	datazone_createEnvironmentBlueprintCmd.MarkFlagRequired("name")
	datazone_createEnvironmentBlueprintCmd.MarkFlagRequired("provisioning-properties")
	datazoneCmd.AddCommand(datazone_createEnvironmentBlueprintCmd)
}
